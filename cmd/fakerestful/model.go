package main

var db [9]user

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
