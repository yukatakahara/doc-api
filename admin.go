package admin

// todo: use context to keep my db connection
// or use a service structure

import (
	"log"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/schema"
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// var dbPath = "C:/Users/Alan/Projects/data/db.boltdb"
var MySigningKey = []byte("secret")

func init() {
	schema.RegisterType("Admin", Admin{})
	schema.RegisterType("Clinic", Clinic{})
	schema.GenerateID = func(_ interface{}) quad.Value {
		return quad.IRI(uuid.NewV1().String())
	}
}

type Admin struct {
	Name           string `json:"name" quad:"name"`
	Email          string `json:"email" quad:"email"`
	HashedPassword string `json:"hashedPassword"  quad:"hashed_password"`
	LoggedIn       bool
	Store          *cayley.Handle
}

type AdminService interface {
	// Admin(id int) (*Admin, error)
	// Admins() ([]*Admin, error)
	CreateAdmin(u *Admin, password string) error
	// DeleteAdmin(id int) error
}

type Clinic struct {
	Name      string   `json:"name" quad:"name"`
	Address1  string   `json:"address" quad:"address"`
	CreatedBy quad.IRI `quad:"createdBy"`
}

type EmailAndPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewClinic struct {
	Name     string `json:"name"`
	Address1 string `json:"address1"`
}

type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func New(store *cayley.Handle) (*Admin, error) {
	a := &Admin{}
	a.Store = store
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
