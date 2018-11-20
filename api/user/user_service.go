package user

import (
	"github.com/kongebra/cpts/api/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User Service deals with database functions
type Service struct {
	Collection *mgo.Collection
}

/*
Creates new service for the User-collection for the database
 */
func NewUserService(session *mongo.Session, dbName string, collectionName string) *Service {
	collection := session.GetCollection(dbName, collectionName)

	for _, key := range Index() {
		collection.EnsureIndex(key)
	}

	return &Service{collection}
}

/*
Create a new user to the database.
 */
func (s *Service) Create(u *User) error {
	return s.Collection.Insert(u)
}

/*
GetAll returns a slice with all the users.
 */
func (s *Service) GetAll() ([]User, error) {
	all := make([]User, 0)
	err := s.Collection.Find(nil).All(&all)
	return all, err
}

/*
GetByUsername returns a user with given username
 */
func (s *Service) GetByUsername(username string) (*User, error) {
	user := User{}
	err := s.Collection.Find(bson.M{"username": username}).One(&user)
	return &user, err
}

/*
GetByID returns a user with given ID
 */
func (s *Service) GetByID(id bson.ObjectId) (*User, error) {
	user := User{}
	err := s.Collection.Find(bson.M{"_id": id}).One(&user)
	return &user, err
}

/*
Update updates the information within the user
 */
func (s *Service) Update(user *User) error {
	return s.Collection.Update(bson.M{"_id": user.Id}, user)
}
