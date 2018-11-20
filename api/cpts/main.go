package cpts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kongebra/cpts/api/middleware"
	"log"
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
	api.Router.Use(middleware.Logger)
	api.loadFromDB()
	api.registerRoutes()

	var err error
	api.Session, err = mongo.NewSession("ds143532.mlab.com:43532")

	if err != nil {
		panic(err)
	}

	api.EventService = event.NewEventService(api.Session, "cpts", "event")
	api.UserService = user.NewUserService(api.Session, "cpts", "user")
	api.TicketService = ticket.NewTicketService(api.Session, "cpts", "ticket")
}

func (api *CPTS) registerRoutes() {

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

	api.Router.HandleFunc("/api/event", func(w http.ResponseWriter, r *http.Request) {
		events, err := api.EventService.GetAll()

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(events)
	}).Methods("GET")

	api.Router.HandleFunc("/api/event/webhooks", func(w http.ResponseWriter, r *http.Request) {
		wh, err := event.RegisterWebhook(w, r)
		if err == nil {
			api.Webhooks = append(api.Webhooks, wh)
		}
	}).Methods("POST")

	api.Router.HandleFunc("/api/ticket", func(w http.ResponseWriter, r *http.Request) {
		tickets, err := api.TicketService.GetAll()

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(tickets)
	}).Methods("GET")

	api.Router.HandleFunc("/api/ticket", func(w http.ResponseWriter, r *http.Request) {
		t := ticket.Ticket{
			Id: bson.NewObjectId(),
			Event: bson.ObjectIdHex(r.FormValue("event")),
			Scanned: false,
		}

		u, err := api.UserService.GetByID(bson.ObjectIdHex(r.FormValue("user")))

		if err != nil {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		e, err := api.EventService.GetById(bson.ObjectIdHex(r.FormValue("event")))

		if err != nil {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		e.Participants = append(e.Participants, t.Id)
		err = api.EventService.Update(e)

		if err != nil {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		u.Tickets = append(u.Tickets, t.Id)
		err = api.UserService.Update(u)

		if err != nil {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = api.TicketService.Create(&t)

		if err != nil {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, t.Id)
	}).Methods("POST")

	api.Router.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		u := user.User{
			Id: bson.NewObjectId(),
			Username: r.FormValue("username"),
			Email: r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		err := api.UserService.Create(&u)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, u.Id)
	}).Methods("POST")

	api.Router.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		users, err := api.UserService.GetAll()

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(users)
	}).Methods("GET")

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
