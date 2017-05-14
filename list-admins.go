package admin

import (
	"log"
	"regexp"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
)

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
