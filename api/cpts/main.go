package cpts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kongebra/cpts/api/mongo"
	"github.com/kongebra/cpts/api/ticket"

	"github.com/gorilla/mux"
	"github.com/kongebra/cpts/api/event"
	"github.com/kongebra/cpts/api/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CPTS struct {
	Users         []user.User `json:"users"`
	Events        []event.Event
	Webhooks      []event.Webhook
	Router        *mux.Router
	Session       *mongo.Session
	EventService  *event.Service
	TicketService *ticket.Service
	UserService   *user.Service
}

func (api *CPTS) Init() {
	api.Router = mux.NewRouter().StrictSlash(true)
	api.loadFromDB()
	api.registerRoutes()

	var err error
	api.Session, err = mongo.NewSession("ds143532.mlab.com:43532")

	if err != nil {
		panic(err)
	}

	api.EventService = event.NewEventService(api.Session, "cpts", "event")
}

func (api *CPTS) AddUser(u user.User) {
	api.Users = append(api.Users, u)
}

func (api *CPTS) registerRoutes() {

	api.Router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hei")
	})

	api.Router.HandleFunc("/api/event", func(w http.ResponseWriter, r *http.Request) {
		evt, err := event.AddEventHandler(w, r, api.EventService)
		if err == nil {
			api.Events = append(api.Events, evt)

			resp := event.WebhookResponse{}
			resp.Content = "New Event"
			resp.Embeds = append(resp.Embeds, event.WebhookEmbed{})
			resp.Embeds[0].Title = evt.Name
			resp.Embeds[0].Image.URL = evt.IMGURL

			respBytes, marshalErr := json.Marshal(resp)
			if marshalErr != nil {
				return
			}

			for _, v := range api.Webhooks {
				http.Post(v.URL, "application/json", bytes.NewBuffer(respBytes))
			}
		}
	}).Methods("POST")

	api.Router.HandleFunc("/api/event/webhooks", func(w http.ResponseWriter, r *http.Request) {
		wh, err := event.RegisterWebhook(w, r)
		if err == nil {
			api.Webhooks = append(api.Webhooks, wh)
		}
	}).Methods("POST")

}

func (api *CPTS) loadFromDB() {
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
		println("failed to read from db")
		return
	}

	errEvt := session.DB("cpts").C("events").Find(bson.M{}).All(&api.Events)
	if errEvt != nil {
		println("failed to read from events collection")
		return
	}

	errWh := session.DB("cpts").C("webhooks").Find(bson.M{}).All(&api.Webhooks)
	if errWh != nil {
		println("failed to read from webhooks collection")
		return
	}

}
