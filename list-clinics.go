package admin

import (
	"fmt"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/schema"
)

// get admins from the db
func (a *Admin) AllClinics() ([]Clinic, error) {
	// As, err := ReadAdmins(store, regexp.MustCompile(".*"))
	As, err := readAllClinics(store)
	printAllQuads2(store)
	// printAllClinics(store)

	if err != nil {
		return []Clinic{}, err
	}

	// return As
	return As, nil
}

func printAllQuads2(store *cayley.Handle) {
	it := store.QuadsAllIterator()
	defer it.Close()
	fmt.Println("\nquads:")
	for it.Next() {
		fmt.Println(store.Quad(it.Result()))
	}
	fmt.Println()
}

func readAllClinics(store *cayley.Handle) ([]Clinic, error) {
	// get all clinics
	var clinics []Clinic
	err := schema.LoadTo(nil, store, &clinics)

	return clinics, err
}

func printAllClinics(store *cayley.Handle) {
	// get all clinics
	var clinics []Clinic
	CheckErr(schema.LoadTo(nil, store, &clinics))
	fmt.Println("clinics:")
	for _, a := range clinics {
		fmt.Printf("%+v\n", a)
	}
	fmt.Println()
}
