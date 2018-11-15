package cpts

import (
	"github.com/kongebra/cpts/cmd/event"
	"github.com/kongebra/cpts/cmd/user"
)

type CPTS struct {
	Users []user.User `json:"users"`
	Events []event.Event `json:"events"`
}