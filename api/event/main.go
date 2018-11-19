package event

import (
	"github.com/gorilla/mux"
	"github.com/kongebra/cpts/api/ticket"
	"gopkg.in/mgo.v2/bson"
)

type TimeInterval struct {
	Start string `json:"start"`
	End string `json:"end"`
}

type Event struct {
	Id bson.ObjectId `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Date TimeInterval `json:"date"`
	Participants []bson.ObjectId `json:"participants"`
}

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/api/events/", getEvents).Methods("GET")
	r.HandleFunc("/api/events/{id}", getEvent).Methods("GET")
	r.HandleFunc("/api/events", createEvent).Methods("POST")
	r.HandleFunc("/api/events/{id}", updateEvents).Methods("PUT")
	r.HandleFunc("/api/events/{id}", deleteEvents).Methods("DELETE")
}

func (e *Event) CreateTicket() ticket.Ticket {
	t := ticket.Ticket{
		Id: bson.NewObjectId(),
		Event: e.Id,
		Scanned: false,
	}

	return t
}

func (e *Event) ScanTicket(t *ticket.Ticket) bool {
	if t.Scanned {
		return false
	} else {
		if t.Event == e.Id {
			t.Scanned = true
			return true
		} else {
			return false
		}
	}
}