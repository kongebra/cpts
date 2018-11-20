package ticket

import (
	"github.com/kongebra/cpts/api/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	Collection *mgo.Collection
}

func NewTicketService(session *mongo.Session, dbName string, collectionName string) *Service {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(Index())
	return &Service{collection}
}

func (s *Service) Create(t *Ticket) error {
	return s.Collection.Insert(t)
}

func (s *Service) GetAll() ([]Ticket, error) {
	all := make([]Ticket, 0)
	err := s.Collection.Find(bson.M{}).All(&all)
	return all, err
}

func (s *Service) GetByID(id bson.ObjectId) (*Ticket, error) {
	ticket := Ticket{}
	err := s.Collection.Find(bson.M{"_id": id}).One(&ticket)
	return &ticket, err
}

func (s *Service) GetByEventID(id bson.ObjectId) ([]Ticket, error) {
	all := make([]Ticket, 0)
	err := s.Collection.Find(bson.M{"event": id}).All(&all)
	return all, err
}