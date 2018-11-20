package test

import (
	"github.com/kongebra/cpts/api/db"
	"github.com/kongebra/cpts/api/user"
	"gopkg.in/mgo.v2"
	"testing"
)

func TestUserDatabaseInit(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	for _, key := range []string{"username", "email"} {
		index := mgo.Index{
			Key:        []string{key},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}

		err = session.DB(db.DATABASE).C("user_test").EnsureIndex(index)

		if err != nil {
			t.Error("Could not ensure index")
		}
	}
}

func TestUserDatabaseInsert(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	u := user.User{
		Username: "janedoe",
		Email: "jane@doe.com",
		Password: "johndoe123",
	}

	err = session.DB(db.DATABASE).C("user").Insert(u)

	if err != nil {
		t.Error("Could not insert data into collection")
	}
}

func TestUserDatabaseInsertDuplicate(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	u := user.User{
		Username: "kongebra",
		Email: "contact@kongebra.net",
		Password: "golang123",
	}

	err = session.DB(db.DATABASE).C("user").Insert(u)

	if err != nil {
		t.Error("Could not insert data into collection")
	}

	err = session.DB(db.DATABASE).C("user").Insert(u)

	if err != nil {
		t.Error("Could not insert data into collection")
	}
}

