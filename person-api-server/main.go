package main

import(
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  "personApi/prom"
)

type Person struct {
  ID        string `json:"id,omitempty"`
  Firstname string `json:"firstname,omitempty"`
  Lastname  string `json:"lastname,omitempty"`
  Address *Address `json:"address,omitempty"`
}

type Address struct {
  City string `json:"city,omitempty"`
  State string `json:"state,omitempty"`

}

//global variable
var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
  id:= req.URL.Query().Get("id")
  for _, item := range people {
    if item.ID == id {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{})
}


func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
  json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
  var person Person
  _ = json.NewDecoder(req.Body).Decode(&person)
  person.ID = fmt.Sprintf("%d", len(people)+1)
  people = append(people,person)
  json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
  id := req.URL.Query().Get("id")
  for index,item := range people {
    if item.ID == id {
      people = append(people[:index], people[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(people)
}


func main(){

  //mock data
  people = append(people, Person{ID: "1", Firstname:"Shreyas", Lastname:"Gune", Address: &Address{City: "Falls Church", State:"Virginia"}})
  people = append(people, Person{ID: "2", Firstname:"Kenshi", Lastname:"Himura"})



  http.HandleFunc("/people", GetPeopleEndpoint)
  http.HandleFunc("/people/create", CreatePersonEndpoint)
  http.HandleFunc("/people/delete", DeletePersonEndpoint)
  http.HandleFunc("/people/get", GetPersonEndpoint)


  //Prometheus Handler
  prom.MetricsEndpoint()
  http.Handle("/metrics", promhttp.Handler())
  log.Fatal(http.ListenAndServe(":2112", nil))

}
