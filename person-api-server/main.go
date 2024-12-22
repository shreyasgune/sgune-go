// person-api-server is a simple HTTP server for managing people and integrating with Prometheus metrics.
package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"personApi/prom"
)

// Person represents an individual person with an ID, first name, last name, and an optional address.
type Person struct {
	ID        string  `json:"id,omitempty"`
	Firstname string  `json:"firstname,omitempty"`
	Lastname  string  `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address represents a location with a city and state.
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// people is a global variable that stores a list of people.
var people []Person

// GetPersonEndpoint handles the "/people/get" endpoint and returns a person based on the provided ID.
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	for _, item := range people {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// Return an empty person if not found
	json.NewEncoder(w).Encode(&Person{})
}

// GetPeopleEndpoint handles the "/people" endpoint and returns all people.
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// CreatePersonEndpoint handles the "/people/create" endpoint and creates a new person.
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = fmt.Sprintf("%d", len(people)+1)
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePersonEndpoint handles the "/people/delete" endpoint and deletes a person based on the provided ID.
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

// main function initializes mock data, sets up routes, and starts the HTTP server.
func main() {
	// Mock data
	people = append(people, Person{ID: "1", Firstname: "Shreyas", Lastname: "Gune", Address: &Address{City: "Falls Church", State: "Virginia"}})
	people = append(people, Person{ID: "2", Firstname: "Kenshi", Lastname: "Himura"})

	// HTTP Handlers
	http.HandleFunc("/people", GetPeopleEndpoint)
	http.HandleFunc("/people/create", CreatePersonEndpoint)
	http.HandleFunc("/people/delete", DeletePersonEndpoint)
	http.HandleFunc("/people/get", GetPersonEndpoint)

	// Prometheus Metrics
	prom.MetricsEndpoint()
	http.Handle("/metrics", promhttp.Handler())

	// Start the server
	log.Fatal(http.ListenAndServe(":2112", nil))
}