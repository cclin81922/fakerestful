package main

import (
	"log"
	"net/http"
)

func init() {
	for idx := range db {
		db[idx].ID = idx + 1
	}
}

func main() {
	http.HandleFunc("/users/", usersAPIHandler)

	log.Println("RESTful API is serving at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
