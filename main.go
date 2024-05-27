package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/people", peopleHandler)
	log.Println("Server listening to port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}