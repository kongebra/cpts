package event

import (
	"gopkg.in/mgo.v2/bson"
)

type TimeInterval struct {
	Start string `json:"start"`
	End string `json:"end"`
}

type Event struct {
	Id bson.ObjectId `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Date TimeInterval `json:"date"`
	Participants []bson.ObjectId `json:"participants"`
}

