// A stand-alone HTTP server
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Message struct represents the JSON document which the API sends when something wrong will happen.
// Type: If this is an error message, "Error" or "Message" can be the two values.
// UserMessage: This is the string which the developer can show to the user.
// DeveloperMessage: This is the technical message.
// DocumentationLink: If the message type is error, this will be the link to corresponding documentation.
type Message struct {
	Type              string `json:"messageType"`
	UserMessage       string `json:"userMessage"`
	DeveloperMessage  string `json:"developerMessage"`
	DocumentationLink string `json:"documentationLink"`
}

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var mySigningKey = []byte("secret")

func init() {
	// POST /register - create jwt
	http.HandleFunc("/register", GetTokenHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
}

type User struct {
	Email string `json:"email"`
	JWT   string `json:"jwt"`
}

//GetTokenHandler will get a token for the username and password
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
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
		tokenString, err := token.SignedString(mySigningKey)
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

// ReturnMessageJSON is a wrapper which will send JSON document of type Message, it takes the following arguments
// messageType: Error or Information, denotes if the message is just FYI or an error
// userMessage: Message in terms of user
// devMessage: Message in terms of Developer
func ReturnMessageJSON(w http.ResponseWriter, messageType, userMessage, devMessage string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token")
	w.Header().Set("Content-Type", "application/json")

	message := Message{Type: messageType, UserMessage: userMessage, DeveloperMessage: devMessage}
	if messageType == "Information" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		panic(err)
	}
	return
}
