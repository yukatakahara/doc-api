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

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var dbPath = "/tmp/db.boltdb"
var ErrBadFormat = errors.New("invalid email format")
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Admin struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
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

	fmt.Println("email", email)
	err := ValidateFormat(email)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	uuid := uuid.NewV1().String()
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	store.AddQuad(quad.Make(uuid, "is_a", "admin", nil))
	store.AddQuad(quad.Make(uuid, "email", email, nil))
	store.AddQuad(quad.Make(uuid, "hashed_password", hash, nil))

	a := Admin{email, hash}

	return a
}

func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// get admins from the db
func (a *Admin) All() []Admin {
	store := initializeAndOpenGraph(dbPath)
	var err error

	p := cayley.StartPath(store).Has(quad.String("is_a"), quad.String("admin")).Out()

	results := []Admin{}
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		nativeValue := quad.NativeOf(value) // this converts RDF values to normal Go types
		fmt.Println(nativeValue)
	})

	if err != nil {
		log.Fatalln(err)
	}

	results = append(results, Admin{"foo@gmail.com", "mbmbjkjk"})
	results = append(results, Admin{"bar@gmail.com", "54354353"})

	return results
}
