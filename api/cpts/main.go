package cpts

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kongebra/cpts/api/event"
	"github.com/kongebra/cpts/api/middleware"
	"gopkg.in/mgo.v2"
	"net/http"
	"time"
)

var (
	ADDRS = []string{"ds133533.mlab.com:33533"}
	DATABASE = "assidnment-2"
	USERNAME = "cpts"
	PASSWORD = "cpts123"
)

func RouterManager(r *mux.Router) {
	var dialInfo = &mgo.DialInfo{
		Addrs: ADDRS,
		Database: DATABASE,
		Username: USERNAME,
		Password: PASSWORD,
		Timeout: 60 * time.Second,
	}

	r.Use(middleware.Logger)

	r.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer,"/API")
	}).Methods("GET")

	r.HandleFunc("/api/event", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello")
		event.Create(writer, request, dialInfo)
	}).Methods("POST")
}