package mongo

import (
	"time"

	"gopkg.in/mgo.v2"
)

type Session struct {
	session *mgo.Session
}

func NewSession(url string) (*Session, error) {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{url},
		Database: "cpts",
		Username: "test12",
		Password: "test12",
		Timeout:  60 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo)

	return &Session{session}, err
}

func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

func (s *Session) GetCollection(database string, collection string) *mgo.Collection {
	return s.session.DB(database).C(collection)
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

func (s *Session) DropDatabase(db string) error {
	if s.session != nil {
		return s.session.DB(db).DropDatabase()
	}

	return nil
}
