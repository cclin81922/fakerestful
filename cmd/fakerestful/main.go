package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", usersAPIHandler)

	log.Println("RESTful API is serving at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
