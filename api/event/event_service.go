package event

import (
	"github.com/kongebra/cpts/api/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	Collection *mgo.Collection
}

func NewEventService(session *mongo.Session, dbName string, collectionName string) *Service {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(Index())
	return &Service{collection}
}

func (s *Service) Create(e *Event) error {
	return s.Collection.Insert(e)
}

func (s *Service) GetAll() ([]Event, error) {
	all := make([]Event, 0)
	err := s.Collection.Find(nil).All(&all)
	return all, err
}

func (s *Service) GetByName(name string) (*Event, error) {
	event := Event{}
	err := s.Collection.Find(bson.M{"name": name}).One(&event)
	return &event, err
}

func (s *Service) GetById(id bson.ObjectId) (*Event, error) {
	event := Event{}
	err := s.Collection.Find(bson.M{"_id": id}).One(&event)
	return &event, err
}
<<<<<<< HEAD
=======
func (s *Service) Update(event *Event) error {
	return s.Collection.Update(bson.M{"_id": event.Id}, event)
}

>>>>>>> d346240cf6028a29fcf40a4c8ae0fa381410648b
