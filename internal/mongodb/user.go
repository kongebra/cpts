package mongodb

import (
	"github.com/kongebra/cpts/cmd/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type UserDB struct {
	Database Database
	Collection string
}

func (db *UserDB) Dial() (*mgo.Session, error) {
	return db.Database.Dial()
}

func (db *UserDB) Init() {
	session, _ := db.Dial()

	defer session.Close()

	index := mgo.Index{
		Key: []string{"id", "username", "email", "phone"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}

	err := session.DB(db.Database.Database).C(db.Collection).EnsureIndex(index)

	if err != nil {
		log.Fatalf("EnsureIndex: %s\n", err)
	}
}

func (db *UserDB) Insert(u user.User) (bson.ObjectId, error) {
	session, _ := db.Dial()

	defer session.Close()

	id := bson.NewObjectId()

	u.Id = id

	err := session.DB(db.Database.Database).C(db.Collection).Insert(u)

	if err != nil {
		log.Fatalf("InsertUser: %s\n", err)
		return "", err
	}

	return id, nil
}
func (db *UserDB) Count() int {
	session, _ := db.Dial()

	defer session.Close()

	count, err := session.DB(db.Database.Database).C(db.Collection).Count()

	if err != nil {
		log.Fatalf("Count: %s\n", err)
		return -1
	}

	return count
}

func (db *UserDB) Delete(id bson.ObjectId) bool {
	session, _ := db.Dial()
	defer session.Close()

	err := session.DB(db.Database.Database).C(db.Collection).Remove(bson.M{"id": id})

	if err != nil {
		log.Fatalf("Delete: %s\n", err)
		return false
	}

	return true
}

func (db *UserDB) Get(id bson.ObjectId) user.User {
	session, _ := db.Dial()
	defer session.Close()

	user := user.User{}
	err := session.DB(db.Database.Database).C(db.Collection).Find(bson.M{"id": id}).One(&user)

	if err != nil {
		log.Fatalf("Get: %s\n", err)
	}

	return user
}

func (db *UserDB) GetAll() []user.User {
	session, _ := db.Dial()
	defer session.Close()

	users := make([]user.User, 0)
	err := session.DB(db.Database.Database).C(db.Collection).Find(bson.M{}).All(&users)

	if err != nil {
		log.Fatalf("GetAll: %s\n", err)
	}

	return users
}