package user_test

import (
	"github.com/kongebra/cpts/api/mongo"
	"github.com/kongebra/cpts/api/user"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

const (
	URL = "localhost:27017"
	DATABASE = "cpts_test"
	COLLECTION = "user"
)

func Test_UserService(t *testing.T) {
	session, err := mongo.NewSession(URL)

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %s", err)
	}

	defer func() {
		session.DropDatabase(DATABASE)
		session.Close()
	}()

	userService := user.NewUserService(session.Copy(), DATABASE, COLLECTION)

	testUsername := "test_user"
	testPassword := "test_password"
	testEmail := "test@email.com"

	u := user.User{
		Username: testUsername,
		Password: testPassword,
		Email: testEmail,
	}

	err = userService.Create(&u)

	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}

	var results []user.User
	session.GetCollection(DATABASE, COLLECTION).Find(nil).All(&results)

	count := len(results)

	if count != 1 {
		t.Errorf("Incorrect number of results. Expected: '1', Got: '%d'", count)
	}

	if results[0].Username != u.Username {
		t.Errorf("Incorrect Username. Excpected: '%s', Got: '%s'", testUsername, results[0].Username)
	}
}

func Test_UserService_GetByUsername(t *testing.T) {
	session, err := mongo.NewSession(URL)

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %s", err)
	}

	defer func() {
		session.DropDatabase(DATABASE)
		session.Close()
	}()

	userService := user.NewUserService(session.Copy(), DATABASE, COLLECTION)

	testUsername := "test_user"
	testPassword := "test_password"
	testEmail := "test@email.com"

	u := user.User{
		Username: testUsername,
		Password: testPassword,
		Email: testEmail,
	}

	err = userService.Create(&u)

	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}

	us, err := userService.GetByUsername(testUsername)

	if err != nil {
		t.Errorf("Unable to get user by username: %s", err)
	}

	if us.Username != testUsername {
		t.Errorf("Incorrect Username. Expected: %s, Got: %s", testUsername, us.Username)
	}

	if us.Password != testPassword {
		t.Errorf("Incorrect Password. Expected: %s, Got: %s", testPassword, us.Password)
	}

	if us.Email != testEmail {
		t.Errorf("Incorrect Email. Expected: %s, Got: %s", testEmail, us.Email)
	}
}

func Test_UserService_GetById(t *testing.T) {
	session, err := mongo.NewSession(URL)

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %s", err)
	}

	defer func() {
		session.DropDatabase(DATABASE)
		session.Close()
	}()

	userService := user.NewUserService(session.Copy(), DATABASE, COLLECTION)

	testId := bson.NewObjectId()
	testUsername := "test_user"
	testPassword := "test_password"
	testEmail := "test@email.com"

	u := user.User{
		Id: testId,
		Username: testUsername,
		Password: testPassword,
		Email: testEmail,
	}

	err = userService.Create(&u)

	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}

	us, err := userService.GetByID(testId)

	if err != nil {
		t.Errorf("Unable to get user by ID: %s", err)
	}

	if us.Username != testUsername {
		t.Errorf("Incorrect Username. Expected: %s, Got: %s", testUsername, us.Username)
	}

	if us.Password != testPassword {
		t.Errorf("Incorrect Password. Expected: %s, Got: %s", testPassword, us.Password)
	}

	if us.Email != testEmail {
		t.Errorf("Incorrect Email. Expected: %s, Got: %s", testEmail, us.Email)
	}
}
