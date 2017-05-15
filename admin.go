package admin

// todo: use context to keep my db connection
// or use a service structure

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"regexp"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/schema"
)

var dbPath = "/tmp/db.boltdb"
var ErrBadFormat = errors.New("invalid email format")
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var store *cayley.Handle

func init() {
	store = initializeAndOpenGraph(dbPath)
}

type Admin struct {
	ID             quad.IRI `quad:"@id"`
	Name           string   `json:"name" quad:"name"`
	Email          string   `json:"email" quad:"email"`
	HashedPassword string   `json:"hashedPassword"  quad:"hashed_password"`
	LoggedIn       bool
}

type EmailAndPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	schema.RegisterType("Admin", Admin{})
}

func genID() quad.IRI {
	return quad.IRI(fmt.Sprintf("%x", rand.Intn(0xffff)))
}

// type Admin struct {
// 	ID             string `json:"id"`
// 	Email          string `json:"email"`
// 	Password       string `json:"password"`
// 	HashedPassword string `json:"hashedPassword"`
// 	Name           string `json:"name"`
// 	LoggedIn       bool
// }

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func New() (*Admin, error) {
	a := &Admin{}
	a.LoggedIn = false

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
