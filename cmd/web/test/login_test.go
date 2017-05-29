package main

import (
	"bytes"
	"net/http"
	"testing"
)

// POST /adminlogin
// expect 200
func TestLogin(t *testing.T) {
	url := "http://localhost:3000/adminlogin"

	// curl -v -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"email":"vera@gmail.com","password":"112233"}' "http://localhost:3000/adminlogin"

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

	if resp.Status == "200 OK" {
		return
	}

	t.Fatal("POST /adminlogin returned", resp.Status)

	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
}
