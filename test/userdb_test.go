package main

import (
	user2 "github.com/kongebra/cpts/cmd/user"
	"github.com/kongebra/cpts/internal/mongodb"
	"testing"
)

func initDB() mongodb.UserDB {
	return mongodb.UserDB{
		Database:mongodb.Database{
			Addrs:[]string{"ds133533.mlab.com:33533"},
			Database: "assignment-2",
			Username: "golang",
			Password: "golang1",
		},
		Collection: "user",
	}
}

func TestDBInsert(t *testing.T) {
	db := initDB()

	db.Init()

	user := user2.User{
		Username: "kongebra",
		Firstname: "Svein",
		Lastname: "Danielsen",
		Phone: "92474481",
		Email: "sveindani@gmail.com",
	}

	_, err := db.Insert(user)

	if db.Count() < 1 {
		t.Error("Should be more then zero user")
	}

	if err != nil {
		t.Errorf("Error: UserDBInsert: %s", err)
	}
}

func TestDBCount(t *testing.T) {
	db := initDB()
	db.Init()

	count := db.Count()

	if count == -1 {
		t.Error("Count should not return -1")
	}
}

func TestDBDelete(t *testing.T) {
	db := initDB()
	db.Init()

	user := user2.User{
		Username: "johndoe",
		Firstname: "John",
		Lastname: "Doe",
		Phone: "99887766",
		Email: "johndoe@gmail.com",
	}

	id, err := db.Insert(user)

	if err != nil {
		t.Error("Error inserting user")
	}

	if !db.Delete(id) {
		t.Error("Could not delete user")
	}
}

func TestDBGet(t *testing.T) {
	db := initDB()
	db.Init()

	user := user2.User{
		Username: "johndoe",
		Firstname: "John",
		Lastname: "Doe",
		Phone: "99887766",
		Email: "johndoe@gmail.com",
	}

	id, err := db.Insert(user)

	if err != nil {
		t.Error("Error inserting user")
	}

	u := db.Get(id)

	if u.Username != "johndoe" {
		t.Error("Username is not 'johndoe'")
	}

	if u.Firstname != "John" {
		t.Error("Firstname is not 'John'")
	}

	if u.Lastname != "Doe" {
		t.Error("Lastname is not 'Doe'")
	}

	if !db.Delete(id) {
		t.Error("Could not delete user")
	}
}

func TestDBGetAll(t *testing.T) {
	db := initDB()
	db.Init()

	users := db.GetAll()

	if len(users) != db.Count() {
		t.Error("Count of retrieved array and DB-count is not the same")
	}
}