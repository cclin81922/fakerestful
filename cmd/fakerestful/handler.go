//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

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
			db[0].Name = r.FormValue("name")
			fmt.Fprintf(w, "%v", db[0])
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
			fmt.Fprintf(w, "%v", db[idx])
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
			fmt.Fprintf(w, "%v", db[idx])
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]/delete$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodDelete {
			// Delete a user
			idx := userIndex(r.URL.Path)
			db[idx].Name = ""
			fmt.Fprintf(w, "%v", db[idx])
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.NotFound(w, r)
	}
}
