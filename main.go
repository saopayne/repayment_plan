package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/generate-plan", GeneratePlanHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
