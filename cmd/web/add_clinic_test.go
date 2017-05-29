package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/oren/doc-api"
)

// POST /adminlogin
// expect 200
// POST /clinics
// expect 200
func TestAddClinic(t *testing.T) {
	email := "vera@gmail.com"
	password := "112233"
	name := "a nice place"
	address1 := "2323 hesting road, singapore"

	jwt := login(t, email, password)
	addClinic(t, jwt, name, address1)
}

// curl -v -H "Authorization: Bearer verylongtokenstring" -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"name":"a nice place","address1":"2323 hesting road, singapore"}' "http://localhost:3000/clinics"
func addClinic(t *testing.T, jwt string, name string, address1 string) {
	url := "http://localhost:3000/clinics"

	newClinic := &admin.NewClinic{
		Name:     name,
		Address1: address1,
	}

	js, err := json.Marshal(newClinic)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		return
	}

	t.Fatal("POST /clinics returned", resp.Status)
}
