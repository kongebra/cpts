package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/user/{id}/", handle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
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