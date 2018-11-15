package cpts

import "github.com/kongebra/cpts/api/user"

type CPTS struct {
	Users []user.User `json:"users"`
}

func (api *CPTS) AddUser(u user.User) {
	api.Users = append(api.Users, u)
}