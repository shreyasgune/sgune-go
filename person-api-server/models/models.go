package models

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