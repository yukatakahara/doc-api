package admin

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/schema"
)

// get admins from the db
func (a *Admin) All() ([]Admin, error) {
	As, err := readAllAdmins(a.Store)

	if err != nil {
		return []Admin{}, err
	}

	return As, nil
}

func readAllAdmins(store *cayley.Handle) ([]Admin, error) {
	// get all admins
	var admins []Admin
	err := schema.LoadTo(nil, store, &admins)

	return admins, err
}
