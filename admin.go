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

	err := validateFormat(email)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	uuid := uuid.NewV1().String()
	hash, _ := hashPassword(password) // ignore error for the sake of simplicity

	store.AddQuad(quad.Make(uuid, "is_a", "admin", nil))
	store.AddQuad(quad.Make(uuid, "email", email, nil))
	store.AddQuad(quad.Make(uuid, "hashed_password", hash, nil))

	a := Admin{email, hash}

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

	p := cayley.StartPath(store).Has(quad.String("is_a"), quad.String("admin")).Save("email", "email").Save("hashed_password", "pass")

	results := []Admin{}
	err := p.Iterate(nil).TagValues(store, func(tags map[string]quad.Value) {
		results = append(results, Admin{
			quad.NativeOf(tags["email"]).(string),
			quad.NativeOf(tags["pass"]).(string),
		})
	})

	if err != nil {
		log.Fatalln(err)
	}

	return results
}
