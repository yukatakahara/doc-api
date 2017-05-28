package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/oren/doc-api"
)

func memberLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	if r.Method != "POST" {
		ReturnMessageJSON(w, "Error", "Page not available", "GetTokenHandler only accepts a POST")
		return
	}

	// r.ParseForm()
	username := "josh"
	password := "password123"
	// log.Println(r.Form)

	if username == "" || password == "" {
		ReturnMessageJSON(w, "Error", "Invalid Username/Password", "Invalid Username or password in GetTokenHandler")
		return
	}

	// if db.ValidUser(username, password) {
	if true {
		// Create the Claim which expires after EXPIRATION_HOURS hrs, default is 5.
		claims := MyCustomClaims{
			username,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		/* Sign the token with our secret */
		tokenString, err := token.SignedString(admin.MySigningKey)
		if err != nil {
			log.Println("Something went wrong with signing token")
			ReturnMessageJSON(w, "Error", "Authentication Failed", "Authentication Failed")

			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token")
		w.Header().Set("Content-Type", "application/json")

		user := User{"josh@gmail.com", tokenString}

		js, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)

		// w.Write([]byte(tokenString))
	} else {
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Authentication Failed")
	}
}
