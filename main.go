package main

import (
	"charlotte_backend/apis/people_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//router people API
	router.HandleFunc("/api/people/findall", people_api.FindAll).Methods("GET")
	router.HandleFunc("/api/people/search/{keyword}", people_api.Search).Methods("GET")
	router.HandleFunc("/api/people/create", people_api.Create).Methods("POST")
	router.HandleFunc("/api/people/update", people_api.Update).Methods("PUT")
	router.HandleFunc("/api/people/delete/{id}", people_api.Delete).Methods("DELETE")

	//router memberships API

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
