package admin

import (
	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

// TODO: validation
func (a *Admin) Authenticate(jwt string) (*MyCustomClaims, error) {
	claim, err := validateToken(jwt)

	if err != nil {
		return nil, err
	}

	return claim, nil
}

//ValidateToken will validate the token and return the claims
func validateToken(myToken string) (*MyCustomClaims, error) {
	fmt.Println("token", myToken)
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
