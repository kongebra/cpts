package event

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kongebra/cpts/api/ticket"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TimeInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Event struct {
	ID           bson.ObjectId   `bson:"_id" json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Date         TimeInterval    `json:"date"`
	Participants []bson.ObjectId `json:"participants"`
	IMGURL       string          `bson:"img_url" json:"img_url"`
}

func Create(w http.ResponseWriter, r *http.Request) (Event, error) {
	var evt Event

	// decoding post data
	decodeErr := json.NewDecoder(r.Body).Decode(&evt)
	if decodeErr != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return Event{}, decodeErr
	}

	evt.ID = bson.NewObjectId()

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"ds143532.mlab.com:43532"},
		Database: "cpts",
		Username: "test12",
		Password: "test12",
		Timeout:  60 * time.Second,
	}

	session, dbErr := mgo.DialWithInfo(dialInfo)
	defer session.Close()
	if dbErr != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return Event{}, dbErr
	}

	evt.ID = bson.NewObjectId()

	session.DB("cpts").C("events").Insert(evt)

	return evt, nil
}

/*func MakeEventRoutes(r *mux.Router) {
	r.HandleFunc("/api/events/", getEvents).Methods("GET")
	r.HandleFunc("/api/events/{id}", getEvent).Methods("GET")
	r.HandleFunc("/api/events", createEvent).Methods("POST")
	r.HandleFunc("/api/events/{id}", updateEvents).Methods("PUT")
	r.HandleFunc("/api/events/{id}", deleteEvents).Methods("DELETE")
}*/

func getEvents(w http.ResponseWriter, r *http.Request) {

}

func getEvent(w http.ResponseWriter, r *http.Request) {

}

func createEvent(w http.ResponseWriter, r *http.Request) {

}

func updateEvents(w http.ResponseWriter, r *http.Request) {

}

func deleteEvents(w http.ResponseWriter, r *http.Request) {

}

func (e *Event) CreateTicket() ticket.Ticket {
	t := ticket.Ticket{
		Id:      bson.NewObjectId(),
		Event:   e.ID,
		Scanned: false,
	}

	return t
}

func (e *Event) ScanTicket(t *ticket.Ticket) bool {
	if t.Scanned {
		return false
	} else {
		if t.Event == e.ID {
			t.Scanned = true
			return true
		} else {
			return false
		}
	}
}
