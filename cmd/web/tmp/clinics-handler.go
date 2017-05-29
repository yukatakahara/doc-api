package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oren/doc-api"
)

func clinicsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	if r.Method != "POST" {
		ReturnMessageJSON(w, "Error", "Page not available", "GetTokenHandler only accepts a POST")
		return
	}

	// authenticate admin
	// validate clinic input
	// create clinic in bolt
	Admin, err := admin.New(store)
	if err != nil {
		ReturnMessageJSON(w, "Error", "Authentication Failed", fmt.Sprintf("Error in admin login: %s", err))
		return
	}

	tokenHeader, found := r.Header["Authorization"]

	if !found {
		fmt.Println("not found token header")
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - no auth token")
		return
	}

	jwt := tokenHeader[0]
	fmt.Println("jwt", jwt)
	var claims *admin.MyCustomClaims
	claims, err = Admin.Authenticate(jwt[7:])

	if err != nil {
		fmt.Println("Error in Auth", err)
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - error with token")
		return
	}

	newClinic := &admin.NewClinic{}

	if err := json.NewDecoder(r.Body).Decode(newClinic); err != nil {
		fmt.Println("1", err)
		ReturnMessageJSON(w, "Error", "Error with decoding of clinic", "Error with decoding of clinic")
		return
	}

	fmt.Println("newClinic", newClinic)

	if newClinic.Name == "" || newClinic.Address1 == "" {
		fmt.Println("2", err)
		ReturnMessageJSON(w, "Error", "Invalid Clinic fields", "Invalid Clinic fields")
		return
	}

	c := &admin.Clinic{
		Name:     newClinic.Name,
		Address1: newClinic.Address1,
	}
	err = Admin.AddClinic(c, claims.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
