package admin

import (
	"fmt"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/schema"
)

// TODO: validation
func (a *Admin) AddClinic(c *Clinic, jwt string) error {
	claims, err := validateToken(jwt)
	if err != nil {
		return fmt.Errorf("Admin token is invalid")
	}

	fmt.Println("claims", claims)

	if !validateClinicFields(c) {
		return fmt.Errorf("Clinic fields are not valid")
	}

	// get admin.ID from bolt
	// var foundAdmin Admin
	// foundAdmin, err = FindAdmin(store, claim.Email)
	// var id quad.IRI
	id, err := findAdminID(store, claims.Email)

	if err != nil {
		return err
	}

	// add ID to clinic
	err = insert(store, Clinic{
		Name:      c.Name,
		Address1:  c.Address1,
		CreatedBy: id,
	})

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

func insert(store *cayley.Handle, o interface{}) error {
	qw := graph.NewWriter(store)
	defer qw.Close() // don't forget to close a writer; it has some internal buffering
	_, err := schema.WriteAsQuads(qw, o)
	return err
}

func validateClinicFields(c *Clinic) bool {
	if c.Name == "" || c.Address1 == "" {
		return false
	}

	return true
}

//ValidateToken will validate the token and return the claims
// func validateToken(myToken string) (*MyCustomClaims, error) {
// 	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(mySigningKey), nil
// 	})

// 	if err != nil {
// 		log.Println("Invalid token.", token)
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, fmt.Errorf("JWT invalid")
// 	}

// 	claims := token.Claims.(*MyCustomClaims)

// 	return claims, nil
// }
