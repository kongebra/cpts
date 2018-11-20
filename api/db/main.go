package db

import "gopkg.in/mgo.v2"

type MongoDB struct {
	Session *mgo.Session
}

var (
	URL = "localhost:27017"
	DATABASE = "cpts"
)

func (mdb *MongoDB) ucol() *mgo.Collection {

}