package event

import (
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// TimeInterval stores details of starting and ending time of an event
type TimeInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
/*
TimeInterval represents the main presistent data structure.
It is of the form:
{
	"start": <value>,
	"end": <value>
}
 */

 // Event stores details of the events
type Event struct {
	ID           bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Date         TimeInterval    `json:"date"`
	Participants []bson.ObjectId `json:"participants"`
	IMGURL       string          `json:"img_url" bson:"img_url"`
}
/*
Event represents the main presistent data structure.
It is of the form:
{
	"id": <value>,
	"name": <value>,
	"description": <value>,
	"date": {
		"start": <value>,
		"end": <value>
	},
	"participants": []
}
 */

/*
Create index's for the database, to make name field unique
*/
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

	// form data
	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
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