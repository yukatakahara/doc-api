package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/oren/doc-api"
)

// POST /adminlogin
// expect 200
// POST /clinics
// expect 200
func TestLogin(t *testing.T) {
	name := "a nice place"
	address1 := "2323 hesting road, singapore"
	jwt := login(t)
	fmt.Println("jwt", jwt)
	addClinic(t, jwt, name, address1)
}

func login(t *testing.T) string {
	url := "http://localhost:3000/adminlogin"

	var jsonStr = []byte(`{"email":"vera@gmail.com", "password":"112233"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Fatal("POST /adminlogin returned", resp.Status)
		return ""
	}

	user := admin.User{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&user)
	if err != nil {
		t.Fatal("Error in decoding json", err)
		return ""
	}

	return user.JWT
}

// curl -v -H "Authorization: Bearer verylongtokenstring" -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"name":"a nice place","address1":"2323 hesting road, singapore"}' "http://localhost:3000/clinics"
func addClinic(t *testing.T, jwt string, name string, address1 string) {
	url := "http://localhost:3000/clinics"

	var jsonStr = []byte(`{"name":name, "address1":address1}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Content-Type", "application/json")

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
