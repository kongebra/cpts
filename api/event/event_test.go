package event_test

import (
	"testing"

	"github.com/kongebra/cpts/api/event"
	"github.com/kongebra/cpts/api/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	URL        = "localhost:27017"
	DATABASE   = "cpts_test"
	COLLECTION = "event"
)

func Test_EventService(t *testing.T) {
	session, err := mongo.NewSession(URL)

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %s", err)
	}

	defer func() {
		session.DropDatabase(DATABASE)
		session.Close()
	}()

	eventService := event.NewEventService(session.Copy(), DATABASE, COLLECTION)

	testName := "HusetLan v2019"
	testDesc := "Some random description"
	testDate := event.TimeInterval{Start: "2019-03-11T16:00:00Z", End: "2019-03-13T12:00:00Z"}

	e := event.Event{
		Name:        testName,
		Description: testDesc,
		Date:        testDate,
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

func Test_EventService_GetByName(t *testing.T) {
	session, err := mongo.NewSession(URL)

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %s", err)
	}

	defer func() {
		session.DropDatabase(DATABASE)
		session.Close()
	}()

	eventService := event.NewEventService(session.Copy(), DATABASE, COLLECTION)

	testName := "HusetLan v2019"
	testDesc := "Some random description"
	testDate := event.TimeInterval{Start: "2019-03-11T16:00:00Z", End: "2019-03-13T12:00:00Z"}

	e := event.Event{
		Name:        testName,
		Description: testDesc,
		Date:        testDate,
	}

	err = eventService.Create(&e)

	if err != nil {
		t.Errorf("Unable to create event: %s", err)
	}

	ev, err := eventService.GetByName(testName)

	if err != nil {
		t.Errorf("Unable to find event: %s", err)
	}

	if ev.Name != testName {
		t.Errorf("Incorrect Name. Expected: %s, Got: %s", testName, ev.Name)
	}

	if ev.Description != testDesc {
		t.Errorf("Incorrect Description. Expected: %s, Got: %s", testDesc, ev.Description)
	}

	if ev.Date.Start != testDate.Start {
		t.Errorf("Incorrect Date.Start. Expected: %s, Got: %s", testDate.Start, ev.Date.Start)
	}

	if ev.Date.End != testDate.End {
		t.Errorf("Incorrect Date.End. Expected: %s, Got: %s", testDate.End, ev.Date.End)
	}
}

func Test_EventService_GetById(t *testing.T) {
	session, err := mongo.NewSession(URL)

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %s", err)
	}

	defer func() {
		session.DropDatabase(DATABASE)
		session.Close()
	}()

	eventService := event.NewEventService(session.Copy(), DATABASE, COLLECTION)

	testId := bson.NewObjectId()
	testName := "HusetLan v2019"
	testDesc := "Some random description"
	testDate := event.TimeInterval{Start: "2019-03-11T16:00:00Z", End: "2019-03-13T12:00:00Z"}

	e := event.Event{
		ID:          testId,
		Name:        testName,
		Description: testDesc,
		Date:        testDate,
	}

	err = eventService.Create(&e)

	if err != nil {
		t.Errorf("Unable to create event: %s", err)
	}

	ev, err := eventService.GetById(testId)

	if err != nil {
		t.Errorf("Unable to find event: %s", err)
	}

	if ev.ID != testId {
		t.Errorf("Incorrect ID. Expected: %s, Got: %s", testId, ev.ID)
	}

	if ev.Name != testName {
		t.Errorf("Incorrect Name. Expected: %s, Got: %s", testName, ev.Name)
	}

	if ev.Description != testDesc {
		t.Errorf("Incorrect Description. Expected: %s, Got: %s", testDesc, ev.Description)
	}

	if ev.Date.Start != testDate.Start {
		t.Errorf("Incorrect Date.Start. Expected: %s, Got: %s", testDate.Start, ev.Date.Start)
	}

	if ev.Date.End != testDate.End {
		t.Errorf("Incorrect Date.End. Expected: %s, Got: %s", testDate.End, ev.Date.End)
	}
}
