package admin

import (
	"fmt"
	"log"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
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

	// get email of admin from JWT
	fmt.Println("email in claim", claim.Email)
	// get admin.ID from bolt
	// add ID to clinic struct

	err = insert(store, Clinic{
		Name:      c.Name,
		Address1:  c.Address1,
		CreatedBy: "a1",
	})

	if err != nil {
		return err
	}

	return nil
}

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
