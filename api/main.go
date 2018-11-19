package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kongebra/cpts/api/event"
	"github.com/kongebra/cpts/api/util"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	event.MakeEventRoutes(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Golang!")
	})

	log.Fatal(http.ListenAndServe(util.GetPort(), router))
}