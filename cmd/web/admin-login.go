package main

import (
	"encoding/json"
	"net/http"

	"github.com/oren/doc-api"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	if r.Method != "POST" {
		ReturnMessageJSON(w, "Error", "Page not available", "GetTokenHandler only accepts a POST")
		return
	}

	a := &admin.EmailAndPassword{}

	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		ServerError(w, err)
		return
	}

	if a.Email == "" || a.Password == "" {
		ReturnMessageJSON(w, "Error", "Invalid Email/Password", "Invalid Email or password in GetTokenHandler")
		return
	}

	jwt, err := adminService.Login(a.Password, a.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user := admin.User{a.Email, jwt}
	js, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
