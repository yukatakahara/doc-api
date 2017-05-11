package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
	uuid "github.com/satori/go.uuid"
)

var ErrBadFormat = errors.New("invalid email format")
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func main() {
	// Some globally applicable things
	graph.IgnoreMissing = true
	graph.IgnoreDuplicates = true

	dbFile := flag.String("db", "db.boltdb", "BoltDB file")
	email := flag.String("email", "", "email")
	password := flag.String("password", "", "password")
	flag.Parse()

	store := initializeAndOpenGraph(dbFile)

	createAdmin(store, *email, *password) // add quads to the graph
	fmt.Println(*email)

	// countOuts(store, "41234214")
	// lookAtOuts(store, "41234214")
	// lookAtIns(store, "41234214")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func initializeAndOpenGraph(dbFile *string) *cayley.Handle {
	graph.InitQuadStore("bolt", *dbFile, nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", *dbFile, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return store
}

func createAdmin(store *cayley.Handle, email string, password string) {
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

	countOuts(store, uuid)
	lookAtOuts(store, uuid)
}

func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

// countOuts ... well, counts Outs
func countOuts(store *cayley.Handle, to string) {
	p := cayley.StartPath(store, quad.String(to)).Out().Count()
	fmt.Printf("\n\ncountOuts for %s: ", to)
	p.Iterate(nil).EachValue(store, func(v quad.Value) {
		fmt.Printf("%d\n", quad.NativeOf(v))
	})
	fmt.Printf("============================================\n")
}

// lookAtOuts looks at the outbound links from the "to" node
func lookAtOuts(store *cayley.Handle, to string) {
	p := cayley.StartPath(store, quad.String(to)) // start from a single node, but we could start from multiple

	// this gives us a path with all the output predicates from our starting point
	p = p.Tag("subject").OutWithTags([]string{"predicate"}).Tag("object")

	fmt.Printf("\nlookAtOuts: subject (%s) -predicate-> object\n", to)
	fmt.Printf("============================================\n")

	p.Iterate(nil).TagValues(nil, func(m map[string]quad.Value) {
		fmt.Printf("%s `%s`-> %s\n", m["subject"], m["predicate"], m["object"])
	})
}

// lookAtIns looks at the inbound links to the "to" node
func lookAtIns(store *cayley.Handle, to string) {
	fmt.Printf("\nlookAtIns: object <-predicate- subject (%s)\n", to)
	fmt.Printf("=============================================\n")

	cayley.StartPath(store, quad.Raw(to)).Tag("object").InWithTags([]string{"predicate"}).Tag("subject").Iterate(nil).TagValues(nil, func(m map[string]quad.Value) {
		fmt.Printf("%s <-`%s` %s\n", m["object"], m["predicate"], m["subject"])
	})

}
