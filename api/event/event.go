package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kongebra/cpts/api/ticket"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TimeInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Event struct {
	ID           bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Date         TimeInterval    `json:"date"`
	Participants []bson.ObjectId `json:"participants"`
	IMGURL       string          `json:"img_url" bson:"img_url"`
}

func Index() mgo.Index {
	return mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func AddEventHandler(w http.ResponseWriter, r *http.Request, service *Service) (Event, error) {
	var evt Event
	fmt.Println(r.Header.Get("Content-Type"))
	// form data
	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		/*r.ParseForm()
		formErr := schema.NewDecoder().Decode(&evt, r.PostForm)
		if formErr != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return Event{}, formErr
		}*/
		evt.Name = r.FormValue("name")
		evt.Description = r.FormValue("description")
		evt.Date = TimeInterval{Start: r.FormValue("date_start"), End: r.FormValue("date_end")}
		evt.IMGURL = r.FormValue("image")
	// json data
	case "application/json":
		decodeErr := json.NewDecoder(r.Body).Decode(&evt)
		if decodeErr != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return Event{}, decodeErr
		}
	// other data
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return Event{}, errors.New("invalid form type")
	}

	evt.ID = bson.NewObjectId()

	// add event to db
	dbErr := service.Create(&evt)
	if dbErr != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return Event{}, dbErr
	}

	return evt, nil
}

func (e *Event) CreateTicket() *ticket.Ticket {
	t := ticket.Ticket{
		Id:      bson.NewObjectId(),
		Event:   e.ID,
		Scanned: false,
	}

	e.Participants = append(e.Participants, t.Id)

	return &t
}

func (e *Event) ScanTicket(t *ticket.Ticket) bool {
	if t.Scanned {
		return false
	} else {
		if t.Event == e.ID {
			t.Scanned = true
		}

		return t.Scanned
	}
}
