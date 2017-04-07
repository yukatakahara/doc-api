// A stand-alone HTTP server
package main

import (
	"encoding/json"
	"fmt"
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

type Doctor struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	EmployeeId      string `json:"employee_id"`
	TeamName        string `json:"team_name"`
	TeamEmail       string `json:"team_email"`
	TeamEmployees   string `json:"team_employees"`
	IdeaDescription string `json:"idea_description"`
}

// our 'database' - for now it will be in memory but later on we'll save it in postgres
// slice that each of it's elements is the Doctor struct
var doctors = []Doctor{
	{1, "dan", "dan@gmail.com", "123", "cats", "cats@gmail.com", "josh, dan, lea", "instagram but for cats"},
	{2, "laura", "laura@gmail.com", "143", "dogs", "dogs@gmail.com", "laura, josh", "social network for dogs"},
}

func init() {
	// POST /signup - create jwt
	http.HandleFunc("/signup", GetTokenHandler)
	http.HandleFunc("/settings", GetSettingsHandler)
	http.HandleFunc("/doctors", DoctorsHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
}

type User struct {
	Email string `json:"email"`
	JWT   string `json:"jwt"`
}

func DoctorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDoctors(w, r)
	default:
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
	}
}

func getDoctors(w http.ResponseWriter, r *http.Request) {
	// first we build the response
	res := struct {
		Doctors []Doctor
		Errors  []string
	}{
		doctors,
		[]string{""},
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// then we encode it as JSON on the response
	enc := json.NewEncoder(w)
	err := enc.Encode(res)

	// And if encoding fails we log the error
	if err != nil {
		fmt.Errorf("encode response: %v", err)
	}
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

// func GetSettingsHandler(w http.ResponseWriter, r *http.Request) {
// look in the db for email/password
// and return user info
// if not found return message about it
// }

// ProductsHandler handles the products page, returns all the products.
func GetSettingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ok, username := CheckTokenGetUsername(r)
	if !ok {
		ReturnMessageJSON(w, "Error", "You are not authorized to do this, login & try again", "Invalid token")
		return
	}

	if r.Method == "GET" {
		var user User
		// user, err = db.GetUser(username)

		// if err != nil {
		// 	ReturnMessageJSON(w, "Error", "Could't get user", "error")
		// 	log.Println(err)
		// 	return
		// }

		user = User{"josh@gmail.com", username}
		js, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)

	}
}

// CheckTokenGetUsername gets the HTTP request as the argument and returns if the token is valid
// which is passed as a header of the name Token and returns the username of the logged in user.
func CheckTokenGetUsername(r *http.Request) (bool, string) {
	token := r.Header["Token"][0]
	ok, username := ValidateToken(token)
	return ok, username
}

//ValidateToken will validate the token
func ValidateToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err != nil {
		log.Println("Invalid token.", token)
		return false, ""
	}

	claims := token.Claims.(*MyCustomClaims)
	return token.Valid, claims.Username
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
