package admin

import (
	"fmt"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/schema"
)

// get admins from the db
func (a *Admin) All() ([]Admin, error) {
	// As, err := ReadAdmins(store, regexp.MustCompile(".*"))
	As, err := readAllAdmins(store)
	printAllQuads(store)
	// printAllAdmins(store)

	if err != nil {
		return []Admin{}, err
	}

	// return As
	return As, nil
}

// func ReadAdmins(h *cayley.Handle, email *regexp.Regexp) ([]Admin, error) {
// 	p := cayley.StartPath(h).
// 		Out(quad.IRI("email")).Regex(email).In(quad.IRI("email")).Has(quad.IRI("is_a"), quad.String("admin")).
// 		Save(quad.IRI("email"), "email").
// 		Save(quad.IRI("hashed_password"), "hashed_password")

// 	results := []Admin{}
// 	err := p.Iterate(nil).TagValues(nil, func(tags map[string]quad.Value) {
// 		results = append(results, Admin{
// 			Email:          quad.NativeOf(tags["email"]).(string),
// 			HashedPassword: quad.NativeOf(tags["hashed_password"]).(string),
// 		})
// 	})
// 	if err != nil {
// 		return []Admin{}, err
// 	}
// 	return results, nil
// }

func printAllQuads(store *cayley.Handle) {
	it := store.QuadsAllIterator()
	defer it.Close()
	fmt.Println("\nquads:")
	for it.Next() {
		fmt.Println(store.Quad(it.Result()))
	}
	fmt.Println()
}

func readAllAdmins(store *cayley.Handle) ([]Admin, error) {
	// get all admins
	var admins []Admin
	err := schema.LoadTo(nil, store, &admins)

	return admins, err
}

func printAllAdmins(store *cayley.Handle) {
	// get all admins
	var admins []Admin
	CheckErr(schema.LoadTo(nil, store, &admins))
	fmt.Println("admins:")
	for _, a := range admins {
		fmt.Printf("%+v\n", a)
	}
	fmt.Println()
}
