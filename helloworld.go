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
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"

	"google.golang.org/appengine"
)

// [END import]
// [START main_func]

//func main() {
//	opt := option.WithCredentialsFile("./serviceAccountKey.json")
//	app, err := firebase.NewApp(context.Background(), nil, opt)
//	if err != nil {
//		print(err)
//	}
//	_ = app
//
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		if r.URL.Path != "/" {
//			http.NotFound(w, r)
//			return
//		}
//		fmt.Fprint(w, "Hello, World!")
//		client, err := app.Auth(r.Context())
//		if err != nil {
//			log.Fatalf("error getting Auth client: %v\n", err)
//		}
//
//		token, err := client.VerifyIDToken(r.Context(), "eyJhbGciOiJSUzI1NiIsImtpZCI6IjZmOGUxY2IxNTY0MTQ2M2M2ZGYwZjMzMzk0YjAzYzkyZmNjODg5YWMiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vbWRzLXByb2plY3QtMzQ5MzE2IiwiYXVkIjoibWRzLXByb2plY3QtMzQ5MzE2IiwiYXV0aF90aW1lIjoxNjUzMjEwMjQyLCJ1c2VyX2lkIjoiZkU1bjNNRVE1bldZTUV1NVpSeFdrcHhlWE9DMiIsInN1YiI6ImZFNW4zTUVRNW5XWU1FdTVaUnhXa3B4ZVhPQzIiLCJpYXQiOjE2NTMyMTAyNDMsImV4cCI6MTY1MzIxMzg0MywicGhvbmVfbnVtYmVyIjoiKzQwNzQxMTk4NjA2IiwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJwaG9uZSI6WyIrNDA3NDExOTg2MDYiXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwaG9uZSJ9fQ.3T1Qx89PYlIHKe8qhAqOghUeLxnT3T1MuoFQiR8t_xa7WzZ_2mE7OiMzmVK7q3sfW53fOGWBvc7gBavwylQRgCjUq-3hxsZaL-Y6bsMmknJh1laZqzOMokCLetJD9OI5Kk9soy8hjqV-IPdWenHYGWew0_fqqUD-dbHjSm0-AQT0A4KPAefbTYGYkklmT_aogDGP5O_IkZQfe0-33XSiS8Hdat6cNd4GtJ-THgzHGnby2dnEESnKONkM33f3ZacssxIjTxvIfAiCqxrBYrdV7DQwV73MPbH-46RDc9xgh_FtlrfpyJI38x5j5_lXNdD6uyXOlmbEbsC7QqymbC0LeA")
//		if err != nil {
//			log.Fatalf("error verifying ID token: %v\n", err)
//		}
//
//		log.Printf("Verified ID token: %v\n", token)
//	})
//	http.HandleFunc("/event", eventHandler)
//	http.ListenAndServe(":3000", nil)
//	appengine.Main()
//}

// [END main_func]

type Event struct {
	Name string
	Date time.Time
}

func testUser(res http.ResponseWriter, req *http.Request) {
	_, _ = res, req

}

// [START indexHandler]

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

//[END eventHandler]
//
//[END gae_go111_app]
