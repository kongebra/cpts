package ticket

import (
	"github.com/kongebra/cpts/api/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Ticket Service deals with database functions
type Service struct {
	Collection *mgo.Collection
}

/*
NewTicketService creates a new service for tickets
 */
func NewTicketService(session *mongo.Session, dbName string, collectionName string) *Service {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(Index())
	return &Service{collection}
}

/*
Create inserts a ticket into the database
 */
func (s *Service) Create(t *Ticket) error {
	return s.Collection.Insert(t)
}

/*
GetAll returns a slice of all the tickets in the database
 */
func (s *Service) GetAll() ([]Ticket, error) {
	all := make([]Ticket, 0)
	err := s.Collection.Find(bson.M{}).All(&all)
	return all, err
}

/*
GetById returns a ticket with a given ID
 */
func (s *Service) GetByID(id bson.ObjectId) (*Ticket, error) {
	ticket := Ticket{}
	err := s.Collection.Find(bson.M{"_id": id}).One(&ticket)
	return &ticket, err
}

/*
GetByEventID returns a slice of tickets that belongs to a given event
 */
func (s *Service) GetByEventID(id bson.ObjectId) ([]Ticket, error) {
	all := make([]Ticket, 0)
	err := s.Collection.Find(bson.M{"event": id}).All(&all)
	return all, err
}