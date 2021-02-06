package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	people = append( people, Person{ "1", "Santiago", "González",
		&Address{ "Bogotá", "Colombia" } } )
	people = append( people, Person{ "2", "Ryan", "Reyes",
		&Address{ "Dubling", "California" } } )

	//endpoints

	router.HandleFunc( "/people", GetPeopleEndPoint ).Methods( "GET" )
	router.HandleFunc( "/people/{id}", GetPersonEndPoint ).Methods( "GET" )
	router.HandleFunc( "/people/{id}", CreatePersonEndPoint ).Methods( "POST" )
	router.HandleFunc( "/people/{id}", DeletePersonEndPoint ).Methods( "DELETE" )

	log.Fatal( http.ListenAndServe( ":3000", router ) )

}

type Person struct {
	ID string `json:"id,omitempty"`
	FirstName string `json:"fist_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Address *Address `json:"address,ommitempty"`
}

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndPoint( w http.ResponseWriter, r *http.Request ) {
	_ = json.NewEncoder(w).Encode(people)
}

func GetPersonEndPoint( w http.ResponseWriter, r *http.Request ) {
	params := mux.Vars( r )
	for _, item := range people  {
		if item.ID == params[ "id" ]{
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndPoint( w http.ResponseWriter, r *http.Request ) {
	params := mux.Vars( r )
	var person Person
	_ = json.NewDecoder( r.Body ).Decode( &person )

	person.ID = params[ "id" ]

	people = append( people, person )
	_ = json.NewEncoder(w).Encode(person)
}

func DeletePersonEndPoint( w http.ResponseWriter, r *http.Request ) {
	params := mux.Vars( r )
	for index, item := range people {
		if item.ID == params["id"] {
			people = append( people[:index], people[index+1:]... )
			break
		}
	}
	_ = json.NewEncoder(w).Encode(people)
}
