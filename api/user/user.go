package user

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User stores the details of the users
type User struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email string `json:"email" bson:"email"`
	Password string `json:"-" bson:"password"`
	Tickets []bson.ObjectId `json:"tickets"`
}

/*
User represents the main presistent data structure.
It is of the form:
{
	"id": <value>,
	"username: <value>,
	"email": <value>,
	"tickets": []
}
 */

/*
Create index's for the database, to make fields unique
 */
func Index() []mgo.Index {
	var indexes []mgo.Index

	for _, key := range []string{"username", "email"} {
		indexes = append(indexes, mgo.Index{
			Key: []string{key},
			Unique: true,
			DropDups: true,
			Background: true,
			Sparse: true,
		})
	}

	return indexes
}
