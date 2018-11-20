package event

import (
	"github.com/kongebra/cpts/api/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Event Service deals with database functions
type Service struct {
	Collection *mgo.Collection
}

/*
Creates new service for the Event-collection for the database
 */
func NewEventService(session *mongo.Session, dbName string, collectionName string) *Service {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(Index())
	return &Service{collection}
}

/*
Create a new event to the database.
 */
func (s *Service) Create(e *Event) error {
	return s.Collection.Insert(e)
}

/*
GetAll returns a slice with all the events.
 */
func (s *Service) GetAll() ([]Event, error) {
	all := make([]Event, 0)
	err := s.Collection.Find(bson.M{}).All(&all)
	return all, err
}

/*
GetByName returns an event with given name
 */
func (s *Service) GetByName(name string) (*Event, error) {
	event := Event{}
	err := s.Collection.Find(bson.M{"name": name}).One(&event)
	return &event, err
}

/*
GetById returns an event with given ID
 */
func (s *Service) GetById(id bson.ObjectId) (*Event, error) {
	event := Event{}
	err := s.Collection.Find(bson.M{"_id": id}).One(&event)
	return &event, err
}

/*
Update, updates the information within the event in the database
 */
func (s *Service) Update(event *Event) error {
	return s.Collection.Update(bson.M{"_id": event.ID}, event)
}
