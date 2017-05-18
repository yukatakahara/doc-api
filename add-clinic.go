package admin

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/schema"
)

// TODO: validation
func (a *Admin) AddClinic(c *Clinic) error {
	err := validateFields()
	if err != nil {
		return err
	}

	// get email of admin from JWT
	// get admin.ID from bolt
	// add ID to clinic struct

	err = insert(store, Clinic{
		ID:        genID(),
		Name:      c.Name,
		Address1:  c.Address1,
		CreatedBy: "a1",
	})

	if err != nil {
		return err
	}

	return nil
}

func validateFields() error {
	return nil
}

func insert(store *cayley.Handle, o interface{}) error {
	qw := graph.NewWriter(store)
	defer qw.Close() // don't forget to close a writer; it has some internal buffering
	_, err := schema.WriteAsQuads(qw, o)
	return err
}
