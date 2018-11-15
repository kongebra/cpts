package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kongebra/cpts/cmd/event"
	"github.com/kongebra/cpts/cmd/user"
	"github.com/kongebra/cpts/internal/cpts"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/", apiHandler).Methods("GET")
	router.HandleFunc("/api/user/{id}/", handle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var api cpts.CPTS

	api.Users = append(api.Users, user.User{
		Id: bson.NewObjectId(),
		Username: "kongebra",
		Firstname: "Svein",
		Lastname: "Danielsen",
		Email: "sveindani@gmail.com",
		Phone: "92484471",
	})

	api.Users = append(api.Users, user.User{
		Id: bson.NewObjectId(),
		Username: "TotalBug",
		Firstname: "Jakob",
		Lastname: "Fonstad",
		Email: "totalbug@gmail.com",
		Phone: "45763957",
	})

	api.Events = append(api.Events, event.Event{
		Id: bson.NewObjectId(),
		Name: "HusetLAN v2019",
		Description: "Er du klar for Gjøvik's koseligste lan i 2019? Lorem ipsum dolor sit amet!",
		Date: event.TimeInterval{
			Start: "2019-04-05T16:00:00Z",
			End: "2019-04-07T12:00:00Z",
		},
	})

	api.Users[0].BuyTicket(&api.Events[0])
	api.Users[1].BuyTicket(&api.Events[0])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api)
}



func handle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 32)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, id)
}