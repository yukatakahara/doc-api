package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oren/doc-api"
)

func DoctorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDoctors(w, r)
	case "OPTIONS":
		getDoctors(w, r)
	default:
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
	}
}

func getDoctors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in getDoctors")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		ReturnMessageJSON(w, "Information", "", "")
		return
	}

	// lat := r.URL.Query()["lat"]
	// lon := r.URL.Query()["lon"]

	// if lat != nil && lon != nil {
	// 	// order the list of doctors based on proximity
	// 	fmt.Println("URL", r.URL)
	// }

	Admin, err := admin.New()
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

	fmt.Println("in getDoctors 1")

	jwt := tokenHeader[0]
	// if no jwt, return 401 - not authorized
	fmt.Println("in getDoctors 2", jwt)
	var claims *admin.MyCustomClaims
	claims, err = Admin.Authenticate(jwt[7:])
	fmt.Println("in getDoctors 3")
	fmt.Println(claims)

	if err != nil {
		fmt.Println("Error in Auth", err)
		ReturnMessageJSON(w, "Error", "Authentication Failed", "Create Clinic - error with token")
		return
	}

	var clinics []admin.Clinic
	clinics, err = Admin.AllClinics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// then we encode it as JSON on the response
	err = json.NewEncoder(w).Encode(clinics)

	// And if encoding fails we log the error
	if err != nil {
		fmt.Println("2222")
		fmt.Errorf("encode response: %v", err)
	}
}
