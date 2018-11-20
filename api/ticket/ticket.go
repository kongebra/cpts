package ticket

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Ticket stores the details of the tickets
type Ticket struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Event   bson.ObjectId `json:"event"`
	Scanned bool          `json:"scanned"`
}
/*
Ticket represents the main presistent data structure.
It is of the form:
{
	"id": <value>,
	"event: <value>,
	"scanned": <value>
}
*/

/*
Create index's for the database, to make id field unique
 */
func Index() mgo.Index {
	return mgo.Index{
		Key: []string{"_id"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
}