package user

import (
	"github.com/kongebra/cpts/api/event"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	Tickets []bson.ObjectId `json:"tickets"`
}

func (u *User) BuyTicket(e event.Event) {
	t := e.CreateTicket()

	u.Tickets = append(u.Tickets, t.Id)
}