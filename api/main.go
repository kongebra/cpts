package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kongebra/cpts/api/cpts"
	"github.com/kongebra/cpts/api/user"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Golang!")
	})

	router.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		var api cpts.CPTS

		api.AddUser(user.User{
			Id: bson.NewObjectId(),
			Username: "kongebra",
			Email: "sveindani@gmail.com",
			Password: "password123",
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(api)
	})

	log.Fatal(http.ListenAndServe(GetPort(), router))
}

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	return ":" + port
}
