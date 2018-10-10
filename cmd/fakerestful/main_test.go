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
	"testing"
)

func TestMain(t *testing.T) {
	http.HandleFunc("/users/", usersAPIHandler)
	go http.ListenAndServe(":8080", nil)

	testcases := []struct {
		name       string
		method     string
		path       string
		statusCode int
	}{
		{"List all users", http.MethodGet, "/users/", http.StatusOK},
		{"Get a user", http.MethodGet, "/users/1", http.StatusOK},
		{"Delete a user", http.MethodGet, "/users/1/delete", http.StatusMethodNotAllowed},
		{"Update a user", http.MethodGet, "/users/1/update", http.StatusMethodNotAllowed},
		{"Get a HTML form for user creation", http.MethodGet, "/users/new", http.StatusOK},
		{"Get a HTML form for user modification", http.MethodGet, "/users/1/edit", http.StatusOK},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080%s", tc.path)
			switch tc.method {
			case http.MethodGet:
				resp, err := http.Get(url)
				if err != nil {
					t.Fatal(err)
				}
				if resp.StatusCode != tc.statusCode {
					t.Fatalf("expected %d | got %d", tc.statusCode, resp.StatusCode)
				}
			}
		})
	}
}
