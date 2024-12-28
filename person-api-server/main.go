// person-api-server is a simple HTTP server for managing people and integrating with Prometheus metrics.
package main

import (

	"time"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"personApi/prom"
	"personApi/models"
)

// people is a global variable that stores a list of people.
var people []models.Person

// LogRequest logs the request method and URL path.
func LogRequest(req *http.Request) {
	log.Printf("Method: %s, Path:%s, User-Agent: %s, Time: %s", req.Method, req.URL.Path, req.UserAgent(), time.Now().Format(time.RFC3339))
}

// LogRequestMetrics increments the request count metric based on the provided method
func LogRequestMetrics(req *http.Request) {
	endpoint := req.URL.Path
	method := req.Method
	prom.RequestCount.WithLabelValues(endpoint, method).Inc()
}

// LogRequestDuration logs the request duration metrics based on the provided method
func LogRequestDuration(req *http.Request) {
	endpoint := req.URL.Path
	method := req.Method
	start := time.Now()
	prom.UpdateRequestDuration(endpoint, method, start)
}

// MetricsInvoker keeps code clean by invoking the metrics functions
func MetricsInvoker(req *http.Request) {
	LogRequest(req)
	LogRequestDuration(req)
	LogRequestMetrics(req)
}

// GetPersonEndpoint handles the "/people/get" endpoint and returns a person based on the provided ID.
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	MetricsInvoker(req)
	found := false
	id := req.URL.Query().Get("id")
	for _, item := range people {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			found = true
			break
		} else {
			json.NewEncoder(w).Encode("Person not found")
			// Return an empty person if not found
			json.NewEncoder(w).Encode(&models.Person{})
			break
		}
	}
	if found == false {
		prom.UpdateErrorCount(req.URL.Path, req.Method, "Person not found")
	}
}

// GetPeopleEndpoint handles the "/people" endpoint and returns all people.
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	MetricsInvoker(req)
	json.NewEncoder(w).Encode(people)
}

// CreatePersonEndpoint handles the "/people/create" endpoint and creates a new person.
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	MetricsInvoker(req)
	var person models.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = fmt.Sprintf("%d", len(people)+1)
	people = append(people, person)
	prom.UpdatePeopleCount(people)
	json.NewEncoder(w).Encode(people)
}

// DeletePersonEndpoint handles the "/people/delete" endpoint and deletes a person based on the provided ID.
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	MetricsInvoker(req)
	id := req.URL.Query().Get("id")
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	prom.UpdatePeopleCount(people)
	json.NewEncoder(w).Encode(people)
}


// UpdatePersonEndpoint handles the "/people/update" endpoint and updates a person based on the provided ID
func UpdatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
MetricsInvoker(req)
	id := req.URL.Query().Get("id")
	var person models.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	for index, item := range people {
		if item.ID == id {
			people[index].Firstname = person.Firstname
			people[index].Lastname = person.Lastname
			people[index].Address = person.Address
			break
		}
	}
	prom.UpdatePeopleCount(people)
	json.NewEncoder(w).Encode(people)
}

// PreBake is a function that initializes mock data and sets up logging.
func PreBake() {

	// HandleLogFormat
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Mock data
	people = append(people, models.Person{ID: "1", Firstname: "Shreyas", Lastname: "Gune", Address: &models.Address{City: "Falls Church", State: "Virginia"}})
	people = append(people, models.Person{ID: "2", Firstname: "Kenshi", Lastname: "Himura"})

	// HTTP Handlers
	http.HandleFunc("/people", GetPeopleEndpoint)
	http.HandleFunc("/people/create", CreatePersonEndpoint)
	http.HandleFunc("/people/delete", DeletePersonEndpoint)
	http.HandleFunc("/people/update", UpdatePersonEndpoint)
	http.HandleFunc("/people/get", GetPersonEndpoint)

	// Prometheus Metrics
	prom.OpsProcessed()
	http.Handle("/metrics", promhttp.Handler())

}

// main function initializes mock data, sets up routes, and starts the HTTP server.
func main() {
	
	PreBake()
	
	// Start the server
	log.Println("Starting server on port 2112")
	log.Fatal(http.ListenAndServe(":2112", nil))
}