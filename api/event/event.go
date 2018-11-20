package event

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TimeInterval stores details of starting and ending time of an event
type TimeInterval struct {
	Start string `json:"start"`
	End string `json:"end"`
}
/*
TimeInterval represents the main presistent data structure.
It is of the form:
{
	"start": <value>,
	"end": <value>
}
 */

 // Event stores details of the events
type Event struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string `json:"name"`
	Description string `json:"description"`
	Date TimeInterval `json:"date"`
	Participants []bson.ObjectId `json:"participants"`
}
/*
Event represents the main presistent data structure.
It is of the form:
{
	"id": <value>,
	"name": <value>,
	"description": <value>,
	"date": {
		"start": <value>,
		"end": <value>
	},
	"participants": []
}
 */

/*
Create index's for the database, to make name field unique
*/
func Index() mgo.Index {
	return mgo.Index{
		Key: []string{"name"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
}