package user

import (
	"github.com/kongebra/cpts/cmd/ticket"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `json:"id"`
	Username string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Phone string `json:"phone"`
	Tickets []ticket.Ticket `json:"tickets"`
}
