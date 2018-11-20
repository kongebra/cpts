package main

import (
	"github.com/gorilla/mux"
	"github.com/kongebra/cpts/api/cpts"
	"github.com/kongebra/cpts/api/util"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	cpts.RouterManager(router)

	log.Fatal(http.ListenAndServe(util.GetPort(), router))
}