// A stand-alone HTTP server
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func init() {
	// GET /projects/  - list of projects
	// POST /projects/ - create project
	http.HandleFunc("/projects/", ProjectsHandler)
	http.HandleFunc("/register", RegisterHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
}

type Project struct {
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
// slice that each of it's elements is the Project struct
var projects = []Project{
	{1, "dan", "dan@gmail.com", "123", "cats", "cats@gmail.com", "josh, dan, lea", "instagram but for cats"},
	{2, "laura", "laura@gmail.com", "143", "dogs", "dogs@gmail.com", "laura, josh", "social network for dogs"},
}

type Profile struct {
	Name    string
	Hobbies []string
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		addProject(w, r)
	case "GET":
		getProjects(w, r)
	default:
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
	}
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	// first we build the response
	res := struct {
		Projects []Project
		Errors   []string
	}{
		projects,
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

func addProject(w http.ResponseWriter, r *http.Request) {
	// decode
	// validate
	// add
	// return to client

	p := &Project{}

	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		ServerError(w, err)
		return
	}

	if err := validateProject(p); err != nil {
		BadRequest(w, err)
		return
	}

	if err := saveProject(&projects, p); err != nil {
		BadRequest(w, errors.New("save error: "+err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// returns the new project - {"id":123,"name":""...}
	json.NewEncoder(w).Encode(p)
}

// we should not use structs to validate a project since as soon as we'll have non-required fields,
// the field appear as empty string. we should use map instead of struct.
// m := make(map[string]string)  _, prs := m["Name"] => prs will be false if Name doesn't exist
func validateProject(p *Project) error {
	if p.Name == "" {
		return fmt.Errorf("missing required fields: %s", "Name")
	}

	return nil
}

// save project to slice. eventualy it will save to a database
func saveProject(projects *[]Project, p *Project) error {
	*projects = append(*projects, *p)

	return nil
}

func BadRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func ServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
