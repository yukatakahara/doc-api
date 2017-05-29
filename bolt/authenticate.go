package bolt

import (
	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/oren/doc-api"
)

func (a *AdminService) Authenticate(jwt string) (*admin.MyCustomClaims, error) {
	claim, err := validateToken(jwt)

	if err != nil {
		return &admin.MyCustomClaims{}, err
	}

	return claim, nil
}

//ValidateToken will validate the token and return the claims
func validateToken(myToken string) (*admin.MyCustomClaims, error) {
	fmt.Println("token", myToken)
	token, err := jwt.ParseWithClaims(myToken, &admin.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(admin.MySigningKey), nil
	})

	if err != nil {
		log.Println("Invalid token.", token)
		return &admin.MyCustomClaims{}, err
	}

	if !token.Valid {
		return &admin.MyCustomClaims{}, fmt.Errorf("JWT invalid")
	}

	claims := token.Claims.(*admin.MyCustomClaims)
	return claims, nil
}
