package bolt

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/schema"
	"github.com/oren/doc-api"
)

func (a *AdminService) All() ([]admin.Admin, error) {
	As, err := readAllAdmins(a.Store)

	if err != nil {
		return []admin.Admin{}, err
	}

	return As, nil
}

func readAllAdmins(store *cayley.Handle) ([]admin.Admin, error) {
	// get all admins
	var admins []admin.Admin
	err := schema.LoadTo(nil, store, &admins)

	return admins, err
}
