package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"

	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/oren/doc-api"
)

var ErrBadFormat = errors.New("invalid email format")
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func main() {
	Admin, err := admin.New()
	if err != nil {
		panic(err)
	}

	results := Admin.All()
	data, err := json.Marshal(results)

	if err != nil {
		fmt.Errorf("encode response: %v", err)
	}

	os.Stdout.Write(data)

	// results := Admin.Create("foobar@gmail.com", "password11")
	// // results := Admin.All()
	// data, err := json.Marshal(results)
	// if err != nil {
	// 	fmt.Errorf("encode response: %v", err)
	// }
	// os.Stdout.Write(data)

	// Some globally applicable things
	// graph.IgnoreMissing = true
	// graph.IgnoreDuplicates = true

	// dbFile := flag.String("db", "db.boltdb", "BoltDB file")
	// email := flag.String("email", "", "email")
	// password := flag.String("password", "", "password")
	// flag.Parse()

	// store := initializeAndOpenGraph(dbFile)

	// createAdmin(store, *email, *password) // add quads to the graph
	// listAdmins(store)
}
