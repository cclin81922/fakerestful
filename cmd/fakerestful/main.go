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

// Command fakerestful is for faking a RESTful API for user resource.
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
