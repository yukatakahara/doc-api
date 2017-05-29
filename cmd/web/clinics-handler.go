package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	tokenHeader, found := r.Header["Authorization"]

	if !found {
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - no auth token")
		return
	}

	jwt := tokenHeader[0]
	fmt.Println("jwt", jwt)
	// var claims *admin.MyCustomClaims

	claims, err := adminService.Authenticate(jwt[7:])
	if err != nil {
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - error with token")
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	newClinic := &admin.NewClinic{}

	err = json.NewDecoder(r.Body).Decode(newClinic)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newClinic.Name == "" || newClinic.Address1 == "" {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c := &admin.Clinic{
		Name:     newClinic.Name,
		Address1: newClinic.Address1,
	}

	err = adminService.AddClinic(c, claims.Email)
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
