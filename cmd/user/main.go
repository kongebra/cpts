package user

import (
	"github.com/kongebra/cpts/cmd/event"
	"github.com/kongebra/cpts/cmd/ticket"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"-" bson:"password"`
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname string `json:"lastname" bson:"lastname"`
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"`
	Tickets []ticket.Ticket `json:"tickets" bson:"tickets"`
}

func (u *User) BuyTicket(event *event.Event) {
	var ticket ticket.Ticket
	ticket.Id = bson.NewObjectId()
	ticket.Event = event.Id

	u.Tickets = append(u.Tickets, ticket)
	event.Participants = append(event.Participants, u.Id)
}

