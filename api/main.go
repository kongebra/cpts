package main

import (
	"log"
	"net/http"

	"github.com/kongebra/cpts/api/cpts"
)

func main() {
	app := cpts.CPTS{}
	app.Init()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
