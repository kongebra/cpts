package user

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id bson.ObjectId `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
}