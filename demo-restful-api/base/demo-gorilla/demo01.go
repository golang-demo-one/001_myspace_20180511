package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	https://blog.csdn.net/wangshubo1989/article/details/71128972?locationNum=8&fps=1
*/

type Person struct {
	ID        string   `json:"id,omitemty"`
	Firstname string   `json:"firstname,omitemty"`
	Lastname  string   `json:"lastname,omitemty"`
	Address   *Address `json:"address,omitemty"`
}

type Address struct {
	City     string `json:"city,omitemty"`
	Province string `json:"province,omitemty"`
}

var people []Person

func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

func GetPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func PostPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "xi", Lastname: "dada", Address: &Address{City: "Shenyang", Province: "Liaoning"}})
	people = append(people, Person{ID: "2", Firstname: "li", Lastname: "xiansheng", Address: &Address{City: "Changchun", Province: "Jilin"}})
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", PostPerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}