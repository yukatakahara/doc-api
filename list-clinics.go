package admin

import (
	"fmt"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/schema"
)

// get all clinics from the db
func (a *Admin) AllClinics() ([]Clinic, error) {
	As, err := readAllClinics(store)

	if err != nil {
		fmt.Println("2222")
		return []Clinic{}, err
	}

	return As, nil
}

func readAllClinics(store *cayley.Handle) ([]Clinic, error) {
	var clinics []Clinic
	err := schema.LoadTo(nil, store, &clinics)

	return clinics, err
}
