package main

import (
	"encoding/json"
	"errors"
	"flag"
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

	email := flag.String("email", "", "email")
	password := flag.String("password", "", "password")
	flag.Parse()

	results := Admin.Create(*email, *password)
	data, err := json.Marshal(results)

	if err != nil {
		fmt.Errorf("encode response: %v", err)
	}

	os.Stdout.Write(data)
}
