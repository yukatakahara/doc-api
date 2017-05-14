package admin

import (
	"errors"
	"log"
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
	ID             string `json:"id"`
	Email          string `json:"email"`
	Password       string `json:"-"` // - so it doesn't get encoded to json ever
	HashedPassword string `json:"hashedPassword"`
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

// TODO: Check for duplicate email
// TODO: Use lock to make sure between check and write we don't have one slip in
// func CreateAdmin(h *cayley.Handle, a Admin) error {
func (a *Admin) Create() error {
	h := initializeAndOpenGraph(dbPath)
	err := validateEmail(a.Email)
	if err != nil {
		return err
	}

	a.HashedPassword, err = hashPassword(a.Password)
	if err != nil {
		return err
	}

	uuid := uuid.NewV1().String()

	t := cayley.NewTransaction()
	t.AddQuad(quad.Make(quad.IRI(uuid), quad.IRI("is_a"), quad.String("admin"), nil))
	t.AddQuad(quad.Make(quad.IRI(uuid), quad.IRI("email"), quad.String(a.Email), nil))
	t.AddQuad(quad.Make(quad.IRI(uuid), quad.IRI("hashed_password"), quad.String(a.HashedPassword), nil))
	err = h.ApplyTransaction(t)

	if err != nil {
		return err
	}

	return nil
}

func validateEmail(email string) error {
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
	h := initializeAndOpenGraph(dbPath)

	As, err := ReadAdmins(h, regexp.MustCompile(".*"))

	if err != nil {
		log.Fatal(err)
	}

	return As

}

func ReadAdmins(h *cayley.Handle, email *regexp.Regexp) ([]Admin, error) {
	p := cayley.StartPath(h).Tag("id").
		Out(quad.IRI("email")).Regex(email).In(quad.IRI("email")).Has(quad.IRI("is_a"), quad.String("admin")).
		Save(quad.IRI("email"), "email").
		Save(quad.IRI("hashed_password"), "hashed_password")

	results := []Admin{}
	err := p.Iterate(nil).TagValues(nil, func(tags map[string]quad.Value) {
		results = append(results, Admin{
			ID:             quad.NativeOf(tags["id"]).(quad.IRI).String(),
			Email:          quad.NativeOf(tags["email"]).(string),
			HashedPassword: quad.NativeOf(tags["hashed_password"]).(string),
		})
	})
	if err != nil {
		return []Admin{}, err
	}
	return results, nil
}
