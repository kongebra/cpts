package ticket

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Ticket struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Event   bson.ObjectId `json:"event"`
	Scanned bool          `json:"scanned"`
}

func Index() mgo.Index {
	return mgo.Index{
		Key: []string{"_id"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
}