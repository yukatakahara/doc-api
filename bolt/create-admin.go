package bolt

import (
	"errors"
	"regexp"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/schema"
	"github.com/oren/doc-api"
	"golang.org/x/crypto/bcrypt"
)

func (a *AdminService) CreateAdmin(newAdmin *admin.Admin, password string) error {
	err := validateEmail(newAdmin.Email)
	if err != nil {
		return err
	}

	var hashedPassword string
	hashedPassword, err = hashPassword(password)
	if err != nil {
		return err
	}

	newAdmin.HashedPassword = hashedPassword

	err = insert(a.Store, newAdmin)

	if err != nil {
		return err
	}

	return nil
}

func validateEmail(email string) error {
	var ErrBadFormat = errors.New("invalid email format")
	var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func insert(store *cayley.Handle, o interface{}) error {
	qw := graph.NewWriter(store)
	defer qw.Close() // don't forget to close a writer; it has some internal buffering
	_, err := schema.WriteAsQuads(qw, o)
	return err
}
