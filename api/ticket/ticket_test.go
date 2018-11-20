package ticket_test

import (
	"github.com/kongebra/cpts/api/event"
	"github.com/kongebra/cpts/api/mongo"
	"github.com/kongebra/cpts/api/ticket"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

const (
	URL = "localhost:27017"
	DATABASE = "cpts_test"
	COLLECTION = "ticket"
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

	eventService := event.NewEventService(session.Copy(), DATABASE, "event")
	ticketService := ticket.NewTicketService(session.Copy(), DATABASE, COLLECTION)

	e := event.Event{
		Name: "Mockup Evnet",
		Description: "Some description",
		Date: event.TimeInterval{
			Start: "13:37",
			End: "04:20",
		},
	}

	err = eventService.Create(&e)

	if err != nil {
		t.Errorf("Unable to create event: %s", err)
	}

	var events []event.Event
	events, err = eventService.GetAll()

	testEvent := events[0].Id
	testScan := false

	tick := ticket.Ticket{
		Id: bson.NewObjectId(),
		Event: testEvent,
		Scanned: testScan,
	}

	err = ticketService.Create(&tick)

	if err != nil {
		t.Errorf("Unable to create ticket: %s", err)
	}

	results, err := ticketService.GetAll()

	if err != nil {
		t.Errorf("Unable to get all Tickets: %s", err)
	}

	count := len(results)

	if count != 1 {
		t.Errorf("Incorrect number of results. Expected: '1', Got: '%d'", count)
	}

	if results[0].Event != tick.Event {
		t.Errorf("Incorrect Username. Excpected: '%s', Got: '%s'", testEvent, results[0].Event)
	}
}

func Test_TicketService_GetByEventId(t *testing.T) {
	session, err := mongo.NewSession(URL)

	if err != nil {
		t.Errorf("Unable to connect to MongoDB: %s", err)
	}

	defer func() {
		session.DropDatabase(DATABASE)
		session.Close()
	}()

	eventService := event.NewEventService(session.Copy(), DATABASE, "event")
	ticketService := ticket.NewTicketService(session.Copy(), DATABASE, COLLECTION)

	e := event.Event{
		Name: "Mockup Evnet",
		Description: "Some description",
		Date: event.TimeInterval{
			Start: "13:37",
			End: "04:20",
		},
	}

	err = eventService.Create(&e)

	if err != nil {
		t.Errorf("Unable to create event: %s", err)
	}

	var events []event.Event
	events, err = eventService.GetAll()

	testEvent := events[0].Id

	if len(events) < 1 {
		e := event.Event{
			Name: "Mockup Evnet",
			Description: "Some description",
			Date: event.TimeInterval{
				Start: "13:37",
				End: "04:20",
			},
		}

		eventService.Create(&e)
		testEvent = e.Id
	} else {
		testEvent = events[0].Id
	}

	for i := 0; i < 5; i++ {
		ti := ticket.Ticket{
			Id: bson.NewObjectId(),
			Event: testEvent,
			Scanned: false,
		}

		ticketService.Create(&ti)
	}

	tickets, err := ticketService.GetByEventID(testEvent)

	if err != nil {
		t.Errorf("Unable to get tickets by event ID: %s", err)
	}

	count := len(tickets)

	if count != 5 {
		t.Errorf("Incorrect number of results. Excepcted: '5', Got: '%d'", count)
	}

	if tickets[0].Event != testEvent {
		t.Errorf("Incorrect event id. Expected: '%s', Got: '%s'", testEvent, tickets[0].Event)
	}
}