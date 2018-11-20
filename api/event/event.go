package event

import (
	"encoding/json"
	"fmt"
	"github.com/kongebra/cpts/api/db"
	"github.com/kongebra/cpts/api/ticket"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
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


var (
	COLLECTION = "events"
)

func InitDB() (*mgo.Session, error) {
	session, err := mgo.Dial(db.URL)

	if err != nil {
		panic(err)
	}

	index := mgo.Index{
		Key: []string{"id", "name"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}

	err = session.DB(db.DATABASE).C(COLLECTION).EnsureIndex(index)

	if err != nil {
		panic(err)
	}

	return session, err
}

func Create(w http.ResponseWriter, r *http.Request, dialInfo *mgo.DialInfo) {
	fmt.Println("TEST")
	session, err := InitDB()
	defer session.Close()

	var id = bson.NewObjectId()

	err = session.DB(dialInfo.Database).C(COLLECTION).Insert(Event{
		Id: id,
		Name: r.FormValue("name"),
		Description: r.FormValue("description"),
		Date: TimeInterval{
			Start: r.FormValue("date_start"),
			End: r.FormValue("date_end"),
		},
	})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(id)
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