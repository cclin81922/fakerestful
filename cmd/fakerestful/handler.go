package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func userIndex(path string) (index int) {
	re, _ := regexp.Compile("[1-9]")
	id, _ := strconv.Atoi(re.FindString(path))
	index = id - 1

	return
}

func usersAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users/" {
		if r.Method == http.MethodGet {
			// List all users
			for _, u := range db {
				fmt.Fprintln(w, u)
			}
		} else if r.Method == http.MethodPost {
			// Create a user
			id := 1
			idx := id - 1
			db[idx].Name = r.FormValue("name")
			fmt.Fprintln(w, "created")
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if r.URL.Path == "/users/new" {
		if r.Method == http.MethodGet {
			// Get a HTML form for user creation
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<form>name: <input /><button>Create</button></form>")
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			// Get a user
			idx := userIndex(r.URL.Path)
			u := db[idx]
			fmt.Fprintf(w, "%v", u)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]/edit$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			// Get a HTML form for user modification
			idx := userIndex(r.URL.Path)
			u := db[idx]
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<form><div>id: %v</div>name: <input value='%v'/><button>Update</button></form>", u.ID, u.Name)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]/update$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodPut {
			// Update a user
			idx := userIndex(r.URL.Path)
			db[idx].Name = r.FormValue("name")
			fmt.Fprintln(w, "updated")
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]/delete$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodDelete {
			// Delete a user
			idx := userIndex(r.URL.Path)
			db[idx].Name = ""
			fmt.Fprintln(w, "deleted")
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.NotFound(w, r)
	}
}
