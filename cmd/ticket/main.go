package ticket

import "gopkg.in/mgo.v2/bson"

type Ticket struct {
	Id bson.ObjectId `json:"id" bson:"id"`
	Event bson.ObjectId `json:"event" bson:"event"`
}