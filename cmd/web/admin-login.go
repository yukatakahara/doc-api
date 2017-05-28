package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oren/doc-api"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		ReturnMessageJSON(w, "Error", "Page not available", "GetTokenHandler only accepts a POST")
		return
	}

	a := &admin.EmailAndPassword{}

	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		fmt.Println("here", err)
		ServerError(w, err)
		return
	}

	if a.Email == "" || a.Password == "" {
		ReturnMessageJSON(w, "Error", "Invalid Email/Password", "Invalid Email or password in GetTokenHandler")
		return
	}

	Admin, err := admin.New()
	if err != nil {
		ReturnMessageJSON(w, "Error", "Authentication Failed", fmt.Sprintf("Error in admin login: %s", err))
		return
	}

	Admin.Email = a.Email

	var jwt string

	jwt, err = Admin.Login(a.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user := User{a.Email, jwt}
	js, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
