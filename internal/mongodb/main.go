package mongodb

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

type Database struct {
	Addrs []string
	Database string
	Username string
	Password string
}

func (db *Database) Dial() (*mgo.Session, error) {
	dialInfo := &mgo.DialInfo{
		Addrs: db.Addrs,
		Database: db.Database,
		Username: db.Username,
		Password: db.Password,
		Timeout: 60 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	return session, err
}
