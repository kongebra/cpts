package event_test

import (
	"github.com/kongebra/cpts/api/event"
	"github.com/kongebra/cpts/api/mongo"
	"testing"
)

const (
	URL = "localhost:27017"
	DATABASE = "cpts_test"
	COLLECTION = "event"
)

func Test_EventService(t *testing.T) {
	session, err := mongo.NewSession(URL)

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %s", err)
	}

	defer func() {
		//session.DropDatabase(DATABASE)
		session.Close()
	}()

	eventService := event.NewEventService(session.Copy(), DATABASE, COLLECTION)

	testName := "HusetLan v2019"
	testDesc := "Some random description"
	testDate := event.TimeInterval{Start: "2019-03-11T16:00:00Z", End: "2019-03-13T12:00:00Z"}

	e := event.Event{
		Name: testName,
		Description: testDesc,
		Date: testDate,
	}

	err = eventService.Create(&e)

	if err != nil {
		t.Errorf("Unable to create event: %s", err)
	}

	var results []event.Event
	session.GetCollection(DATABASE, COLLECTION).Find(nil).All(&results)

	count := len(results)

	if count != 1 {
		t.Errorf("Incorrect number of results. Expected: '1', Got: '%d'", count)
	}

	if results[0].Name != e.Name {
		t.Errorf("Incorrect Username. Excpected: '%s', Got: '%s'", testName, results[0].Name)
	}
}
