// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START gae_go111_app]

// Sample helloworld is an App Engine app.
package main

// [START import]
import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// [END import]
// [START main_func]

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/event", eventHandler)
	appengine.Main()
}

// [END main_func]

type Event struct {
	Name string
	Date time.Time
}

// [START indexHandler]

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

// [END indexHandler]

// [START eventHandler]

func eventHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	e1 := Event{
		Name: "Test",
		Date: time.Now(),
	}

	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "event", nil), &e1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var e2 Event
	if err = datastore.Get(ctx, key, &e2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Stored and retrieved the Event named %q", e2.Name)
}

// [END eventHandler]

// [END gae_go111_app]
