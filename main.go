package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/generate-plan", GeneratePlanHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
