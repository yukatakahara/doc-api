package admin

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// TODO: Check for duplicate email
// TODO: Use lock to make sure between check and write we don't have one slip in
// func CreateAdmin(h *cayley.Handle, a Admin) error {
func (a *Admin) Create() error {
	h := initializeAndOpenGraph(dbPath)
	err := validateEmail(a.Email)
	if err != nil {
		return err
	}

	a.HashedPassword, err = hashPassword(a.Password)
	if err != nil {
		return err
	}

	uuid := uuid.NewV1().String()

	t := cayley.NewTransaction()
	t.AddQuad(quad.Make(quad.IRI(uuid), quad.IRI("is_a"), quad.String("admin"), nil))
	t.AddQuad(quad.Make(quad.IRI(uuid), quad.IRI("email"), quad.String(a.Email), nil))
	t.AddQuad(quad.Make(quad.IRI(uuid), quad.IRI("hashed_password"), quad.String(a.HashedPassword), nil))
	err = h.ApplyTransaction(t)

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
