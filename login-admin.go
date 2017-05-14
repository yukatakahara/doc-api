package admin

import (
	"log"
	"regexp"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"

	"golang.org/x/crypto/bcrypt"
)

func (a *Admin) Login() {
	// find admin in the db based on email
	found, err := findAdmin(store, regexp.MustCompile(a.Email))

	if err != nil {
		log.Fatal(err)
	}

	tmp := checkPasswordHash(a.Password, found.HashedPassword)
	a.LoggedIn = tmp
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func findAdmin(h *cayley.Handle, email *regexp.Regexp) (Admin, error) {
	// defer h.Close()

	p := cayley.StartPath(h).
		Out(quad.IRI("email")).Regex(email).In(quad.IRI("email")).Has(quad.IRI("is_a"), quad.String("admin")).
		Save(quad.IRI("hashed_password"), "hashed_password")

	results := Admin{}
	err := p.Iterate(nil).TagValues(nil, func(tags map[string]quad.Value) {
		results = Admin{
			HashedPassword: quad.NativeOf(tags["hashed_password"]).(string),
		}
	})

	if err != nil {
		return Admin{}, err
	}

	return results, nil
}
