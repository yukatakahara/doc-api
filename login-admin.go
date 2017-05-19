package admin

import (
	"fmt"
	"time"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph/path"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/schema"
	jwt "github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

func (a *Admin) Login(password string) (string, error) {
	// find admin in the db based on email
	adminFound, err := findAdmin(store, a.Email)

	if err != nil {
		return "", err
	}

	passwordMatched := checkPasswordHash(password, adminFound.HashedPassword)
	if !passwordMatched {
		return "", fmt.Errorf("Password incorrect")
	}

	jwt, err := generateJWT(adminFound.Email)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func findAdmin(store *cayley.Handle, email string) (Admin, error) {
	var a Admin
	p := path.StartPath(store).Has(quad.IRI("email"), quad.String(email))
	err := schema.LoadPathTo(nil, store, &a, p)

	if err != nil {
		return a, err
	}

	return a, nil
}

func generateJWT(email string) (string, error) {
	// Create the Claim which expires after EXPIRATION_HOURS hrs, default is 5.
	claims := MyCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	/* Sign the token with our secret */
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("Error while signing a jwt")
	}

	return tokenString, nil
}
