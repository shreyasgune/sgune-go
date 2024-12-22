# person-api-server

`person-api-server` is a simple HTTP server for managing people and integrating with Prometheus metrics.

## Package main

```go
package main
```

## Imports

```go
import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "personApi/prom"
)
```

## Types

### Person

`Person` represents an individual person with an ID, first name, last name, and an optional address.

```go
type Person struct {
    ID        string  `json:"id,omitempty"`
    Firstname string  `json:"firstname,omitempty"`
    Lastname  string  `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
```

### Address

`Address` represents a location with a city and state.

```go
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}
```

## Variables

### people

`people` is a global variable that stores a list of people.

```go
var people []Person
```

## Functions

### GetPersonEndpoint

`GetPersonEndpoint` handles the `/people/get` endpoint and returns a person based on the provided ID.

```go
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    // ...existing code...
}
```

### GetPeopleEndpoint

`GetPeopleEndpoint` handles the `/people` endpoint and returns all people.

```go
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    // ...existing code...
}
```

### CreatePersonEndpoint

`CreatePersonEndpoint` handles the `/people/create` endpoint and creates a new person.

```go
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    // ...existing code...
}
```

### DeletePersonEndpoint

`DeletePersonEndpoint` handles the `/people/delete` endpoint and deletes a person based on the provided ID.

```go
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    // ...existing code...
}
```

### main

`main` function initializes mock data, sets up routes, and starts the HTTP server.

```go
func main() {
    // ...existing code...
}
```

# Prometheus Integration
```
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp
```

# prom
--
    import "."

Package prom provides function to create prometheus metrics

## Usage

#### func  MetricsEndpoint

```go
func MetricsEndpoint()
```
MetricsEndpoint function creates a counter metric and increments it every 2
seconds
