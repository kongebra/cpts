package event

import (
	"github.com/kongebra/cpts/api/ticket"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TimeInterval struct {
	Start string `json:"start"`
	End string `json:"end"`
}

type Event struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string `json:"name"`
	Description string `json:"description"`
	Date TimeInterval `json:"date"`
	Participants []bson.ObjectId `json:"participants"`
}

func Index() mgo.Index {
	return mgo.Index{
		Key: []string{"name"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
}

func (e *Event) CreateTicket() (*ticket.Ticket) {
	t := ticket.Ticket{
		Id: bson.NewObjectId(),
		Event: e.Id,
		Scanned: false,
	}

	e.Participants = append(e.Participants, t.Id)

	return &t
}

func (e *Event) ScanTicket(t *ticket.Ticket) bool {
	if t.Scanned {
		return false
	} else {
		if t.Event == e.Id {
			t.Scanned = true
		}

		return t.Scanned
	}
}