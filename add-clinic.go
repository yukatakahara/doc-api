package admin

import (
	"fmt"
	"log"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/schema"
	jwt "github.com/dgrijalva/jwt-go"
)

// TODO: validation
func (a *Admin) AddClinic(c *Clinic, jwt string) error {
	err := validateFields()
	if err != nil {
		return err
	}

	claim, err := ValidateToken(jwt)

	if err != nil {
		return err
	}

	// get admin.ID from bolt
	// var foundAdmin Admin
	// foundAdmin, err = FindAdmin(store, claim.Email)
	var id string
	// id, err = getAdminID(store, claim.Email)
	id, err = findID(store, claim.Email)

	if err != nil {
		return err
	}

	fmt.Println("id", id)

	// add ID to clinic
	err = insert(store, Clinic{
		Name:      c.Name,
		Address1:  c.Address1,
		CreatedBy: "10",
	})

	if err != nil {
		return err
	}

	return nil
}

func findID(store *cayley.Handle, email string) (string, error) {
	p := cayley.StartPath(store).Has(quad.IRI("email"), quad.String(email))

	id, err := p.Iterate(nil).FirstValue(nil)
	if err != nil {
		return "", err
	}
	return id.(quad.IRI).String(), nil
}

// func getAdminID(store *cayley.Handle, email string) (string, error) {
// 	p := cayley.StartPath(store).Has(quad.IRI("email"), quad.String(email)).Tag("id")

// 	err := p.Iterate(nil).TagValues(nil, func(tags map[string]quad.Value) {
// 		fmt.Println("quad.NativeOf()", quad.NativeOf(tags["id"]).(quad.IRI).String())
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	return "test", nil
// }

// TODO - validate clinic fields
func validateFields() error {
	return nil
}

func insert(store *cayley.Handle, o interface{}) error {
	qw := graph.NewWriter(store)
	defer qw.Close() // don't forget to close a writer; it has some internal buffering
	_, err := schema.WriteAsQuads(qw, o)
	return err
}

func getEmail(jwt string) (string, error) {
	return "foobar@gmail.com", nil
}

//ValidateToken will validate the token and return the claims
func ValidateToken(myToken string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err != nil {
		log.Println("Invalid token.", token)
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("JWT invalid")
	}

	claims := token.Claims.(*MyCustomClaims)
	return claims, nil

}
