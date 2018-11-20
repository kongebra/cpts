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

func (s *Service) Create(u *Event) error {
	return s.Collection.Insert(u)
}

func (s *Service) GetByName(name string) (*Event, error) {
	event := Event{}
	err := s.Collection.Find(bson.M{"name": name}).One(&event)
	return &event, err
}