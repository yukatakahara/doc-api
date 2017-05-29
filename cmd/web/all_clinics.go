package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oren/doc-api"
)

// authenticate admin
// return all clinics
func allClinics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	tokenHeader, found := r.Header["Authorization"]

	if !found {
		fmt.Println("not found token header")
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - no auth token")
		return
	}

	fmt.Println("in getDoctors 1")

	var err error
	jwt := tokenHeader[0]
	_, err = adminService.Authenticate(jwt[7:])

	if err != nil {
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - error with token")
		return
	}

	var clinics []admin.Clinic
	clinics, err = adminService.AllClinics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(clinics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
