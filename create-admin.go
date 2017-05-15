package admin

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/schema"
	"golang.org/x/crypto/bcrypt"
)

// TODO: Check for duplicate email
// TODO: Use lock to make sure between check and write we don't have one slip in
// func CreateAdmin(h *cayley.Handle, a Admin) error {
func (a *Admin) Create(password string) error {
	err := validateEmail(a.Email)
	if err != nil {
		return err
	}

	a.HashedPassword, err = hashPassword(password)
	if err != nil {
		return err
	}

	// uuid := uuid.NewV1().String()
	// t := cayley.NewTransaction()
	// t.AddQuad(quad.Make(quad.IRI(uuid), quad.IRI("is_a"), quad.String("admin"), nil))
	// t.AddQuad(quad.Make(quad.IRI(uuid), quad.IRI("email"), quad.String(a.Email), nil))
	// t.AddQuad(quad.Make(quad.IRI(uuid), quad.IRI("hashed_password"), quad.String(a.HashedPassword), nil))
	// err = store.ApplyTransaction(t)
	err = Insert(store, Admin{
		ID:             genID(),
		Name:           a.Name,
		Email:          a.Email,
		HashedPassword: a.HashedPassword,
	})

	if err != nil {
		return err
	}

	return nil
}

func validateEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Insert(store *cayley.Handle, o interface{}) error {
	qw := graph.NewWriter(store)
	defer qw.Close() // don't forget to close a writer; it has some internal buffering
	_, err := schema.WriteAsQuads(qw, o)
	return err
}
