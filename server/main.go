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
	// GET /doctors/  - list of doctors
	// POST /doctors/ - create doctor
	http.HandleFunc("/doctors/", DoctorsHandler)

	log.Fatal(http.ListenAndServe(":3001", nil))
}

func main() {
}

type Doctor struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	EmployeeId string `json:"employee_id"`
}

// our 'database' - for now it will be in memory but later on we'll save it in postgres
// slice that each of it's elements is the Doctor struct
var doctors = []Doctor{
	{1, "dan", "dan@gmail.com", "123"},
	{2, "laura", "laura@gmail.com", "143"},
}

func DoctorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		addDoctor(w, r)
	case "GET":
		getDoctors(w, r)
	default:
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
	}
}

func getDoctors(w http.ResponseWriter, r *http.Request) {
	// first we build the response
	res := struct {
		Doctors []Doctor
		Errors  []string
	}{
		doctors,
		[]string{""},
	}

	w.Header().Add("Content-Type", "application/json")
	// then we encode it as JSON on the response
	enc := json.NewEncoder(w)
	err := enc.Encode(res)

	// And if encoding fails we log the error
	if err != nil {
		fmt.Errorf("encode response: %v", err)
	}
}

func addDoctor(w http.ResponseWriter, r *http.Request) {
	// decode
	// validate
	// add
	// return to client

	p := &Doctor{}

	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		ServerError(w, err)
		return
	}

	if err := validateDoctor(p); err != nil {
		BadRequest(w, err)
		return
	}

	if err := saveDoctor(&doctors, p); err != nil {
		BadRequest(w, errors.New("save error: "+err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// returns the new doctor - {"id":123,"name":""...}
	json.NewEncoder(w).Encode(p)
}

// we should not use structs to validate a doctor since as soon as we'll have non-required fields,
// the field appear as empty string. we should use map instead of struct.
// m := make(map[string]string)  _, prs := m["Name"] => prs will be false if Name doesn't exist
func validateDoctor(p *Doctor) error {
	if p.Name == "" {
		return fmt.Errorf("missing required fields: %s", "Name")
	}

	return nil
}

// save doctor to slice. eventualy it will save to a database
func saveDoctor(doctors *[]Doctor, p *Doctor) error {
	*doctors = append(*doctors, *p)

	return nil
}

func BadRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func ServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
