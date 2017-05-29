package bolt

import (
	"fmt"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"github.com/oren/doc-api"
)

func (a *AdminService) AddClinic(clinic *admin.Clinic, email string) error {
	if !validateClinicFields(clinic) {
		return fmt.Errorf("Clinic fields are not valid")
	}

	// get admin.ID from bolt
	// var foundAdmin Admin
	// foundAdmin, err = FindAdmin(store, claim.Email)
	// var id quad.IRI
	id, err := findAdminID(a.Store, email)

	if err != nil {
		return err
	}

	clinic.CreatedBy = id
	err = insert(a.Store, clinic)

	if err != nil {
		return err
	}

	return nil
}

func findAdminID(store *cayley.Handle, email string) (quad.IRI, error) {
	p := cayley.StartPath(store).Has(quad.IRI("email"), quad.String(email))
	id, err := p.Iterate(nil).FirstValue(nil)

	if err != nil {
		return "", err
	}

	return id.(quad.IRI), nil
}

func validateClinicFields(c *admin.Clinic) bool {
	if c.Name == "" || c.Address1 == "" {
		return false
	}

	return true
}
