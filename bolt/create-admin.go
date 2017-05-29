package bolt

import "github.com/oren/doc-api"

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
