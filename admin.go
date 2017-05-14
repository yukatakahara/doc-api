package admin

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/schema"
	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"
)

var dbPath = "/tmp/db.boltdb"
var ErrBadFormat = errors.New("invalid email format")
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// type Admin struct {
// 	Email          string `json:"email"`
// 	HashedPassword string `json:"hashedPassword"`
// }

// <admin_id> <rdf:type> <Admin> triple on Admin type
type Admin struct {
	Email          string   `json:"email" quad:"email"`
	HashedPassword string   `json:"hashedPassword"  quad:"hashed_password"`
	ID             quad.IRI `quad:"@id"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func New() (*Admin, error) {
	a := &Admin{}

	return a, nil
}

func initializeAndOpenGraph(dbFile string) *cayley.Handle {
	graph.InitQuadStore("bolt", dbFile, nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", dbFile, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return store
}

// create new admin in the db
func (p *Admin) Create(email string, password string) Admin {
	store := initializeAndOpenGraph(dbPath)

	err := validateFormat(email)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// schema.GenerateID = func(_ interface{}) quad.Value {
	// 	return quad.IRI(uuid.NewV1().String())
	// }

	qw := graph.NewWriter(store)

	hash, _ := hashPassword(password) // ignore error for the sake of simplicity
	id := quad.IRI(uuid.NewV1().String())

	a := Admin{
		email,
		hash,
		id,
	}

	_, err = schema.WriteAsQuads(qw, a)
	checkErr(err)

	return a
}

func validateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// get admins from the db
func (a *Admin) All() []Admin {
	store := initializeAndOpenGraph(dbPath)
	// will require a <admin_id> <rdf:type> <Admin> triple on Admin type
	schema.RegisterType("Admin", Admin{})

	var results []Admin
	err := schema.LoadTo(nil, store, &results)
	checkErr(err)

	return results
}
