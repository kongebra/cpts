package user

import (
	"github.com/kongebra/cpts/api/event"
	"github.com/kongebra/cpts/api/mongo"
	"github.com/kongebra/cpts/api/ticket"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	Collection *mgo.Collection
}

func NewUserService(session *mongo.Session, dbName string, collectionName string) *Service {
	collection := session.GetCollection(dbName, collectionName)

	for _, key := range Index() {
		collection.EnsureIndex(key)
	}

	return &Service{collection}
}

func (s *Service) Create(u *User) error {
	return s.Collection.Insert(u)
}

func (s *Service) GetByUsername(username string) (*User, error) {
	user := User{}
	err := s.Collection.Find(bson.M{"username": username}).One(&user)
	return &user, err
}

func (s *Service) GetByID(id bson.ObjectId) (*User, error) {
	user := User{}
	err := s.Collection.Find(bson.M{"_id": id}).One(&user)
	return &user, err
}

func (s *Service) BuyTicket(e event.Service) (*ticket.Ticket, error) {
	return nil, nil
}