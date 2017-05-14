package admin

// todo: use context to keep my db connection
// or use a service structure

import (
	"errors"
	"log"
	"regexp"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
)

var dbPath = "/tmp/db.boltdb"
var ErrBadFormat = errors.New("invalid email format")
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var store *cayley.Handle

func init() {
	store = initializeAndOpenGraph(dbPath)
}

type Admin struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	HashedPassword string `json:"hashedPassword"`
	LoggedIn       bool
}

func checkErr(err error) {
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
