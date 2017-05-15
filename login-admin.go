package admin

import (
	"fmt"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph/path"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/schema"

	"golang.org/x/crypto/bcrypt"
)

func (a *Admin) Login(password string) error {
	// find admin in the db based on email
	adminFound, err := findAdmin(store, a.Email)

	if err != nil {
		return err
	}

	passwordMatched := checkPasswordHash(password, adminFound.HashedPassword)
	if !passwordMatched {
		return fmt.Errorf("Password incorrect")
	}

	return nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func findAdmin(store *cayley.Handle, email string) (Admin, error) {
	var a Admin
	p := path.StartPath(store).Has(quad.IRI("email"), quad.String(email))
	err := schema.LoadPathTo(nil, store, &a, p)

	if err != nil {
		return a, err
	}

	return a, nil
}

// func findAdmin(h *cayley.Handle, email *regexp.Regexp) (Admin, error) {
// 	// defer h.Close()

// 	p := cayley.StartPath(h).
// 		Out(quad.IRI("email")).Regex(email).In(quad.IRI("email")).Has(quad.IRI("is_a"), quad.String("admin")).
// 		Save(quad.IRI("hashed_password"), "hashed_password")

// 	results := Admin{}
// 	err := p.Iterate(nil).TagValues(nil, func(tags map[string]quad.Value) {
// 		results = Admin{
// 			HashedPassword: quad.NativeOf(tags["hashed_password"]).(string),
// 		}
// 	})

// 	if err != nil {
// 		return Admin{}, err
// 	}

// 	return results, nil
// }
