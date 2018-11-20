package main

import (
	"github.com/kongebra/cpts/api/util"
	"log"
	"net/http"

	"github.com/kongebra/cpts/api/cpts"
)

func main() {
	app := cpts.CPTS{}
	app.Init()

	log.Fatal(http.ListenAndServe(util.GetPort(), app.Router))
}
