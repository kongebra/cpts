package mongo

import (
	"gopkg.in/mgo.v2"
	"time"
)

// Session keeps the database-connection session
type Session struct {
	session *mgo.Session
}

/*
NewSession creates a new session with the database
 */
func NewSession(url string) (*Session, error) {
	dialInfo := &mgo.DialInfo{
		Addrs: []string{url},
		Database: "cpts",
		Username: "test12",
		Password: "test12",
		Timeout: 60 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo)

	return &Session{session}, err
}

/*
Copy makes a copy of the session
 */
func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

/*
GetCollection returns the give collection in the database
 */
func (s *Session) GetCollection(database string, collection string) *mgo.Collection {
	return s.session.DB(database).C(collection)
}

/**
Close shuts down the connection with the database
 */
func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

/*
DropDatabase cleans out the database of all information, used for testing only
 */
func (s *Session) DropDatabase(db string) error {
	if s.session != nil {
		return s.session.DB(db).DropDatabase()
	}

	return nil
}
