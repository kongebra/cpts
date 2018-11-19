package cpts

import (
	"github.com/gorilla/mux"
	"github.com/kongebra/cpts/api/user"
	"net/http"
)

type CPTS struct {
	Users []user.User `json:"users"`
}

func (api *CPTS) AddUser(u user.User) {
	api.Users = append(api.Users, u)
}

func RouterManager(r *mux.Router) {
	r.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		
	})

	r.HandleFunc("/api/event", func(writer http.ResponseWriter, request *http.Request) {

	})
}