package ticket

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type Ticket struct {
	ID      bson.ObjectId `json:"id" bson:"id"`
	Event   bson.ObjectId `json:"event" bson:"event"`
	Scanned bool          `json:"scanned"`
}

func main() {
	// MongoDB mLab connection

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"ds024778.mlab.com:24778"},
		Database: "cpts",
		Username: "test1",
		Password: "test123",
		Timeout:  60 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	defer session.Close()
	if err != nil {
		panic(err)
	}

	// Router
	r := mux.NewRouter()

	r.HandleFunc("/api/ticket/", getTickets).Methods("GET")
	r.HandleFunc("/api/ticket/{id}", getTicket).Methods("GET")
	r.HandleFunc("/api/ticket/{id}", createTicket).Methods("POST") // Event ID
	r.HandleFunc("/api/ticket/{id}", updateTicket).Methods("PUT")
	r.HandleFunc("/api/ticket/{id}", deleteTicket).Methods("DELETE")
}

// Get all tickets
func getTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//TODO Get all Tickets from DB

	json.NewEncoder(w).Encode( /* data  */ ) // TODO Add tickets structure here
}

// Get ticket by param "id"
func getTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params

	//TODO Get all tickets from DB

	// Find specific Ticket by ID

	for _, item := range tickets { // Tickets must be data struct with all tickets
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&tickets{}) //TODO Add tickets from DB
}

// Create Ticket POST Request
func createTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ticket Ticket

	params := mux.Vars(r)

	//TODO Get all events to match id.

	for _, item := range events{
		if item.ID == params["id"]{
			//TODO Add ticket to db
			_ = json.NewDecoder(r.Body).Decode(&ticket)

			ticket.Id = bson.NewObjectId()
			return
		}
	}
	//TODO Throw error (Event not in DB)


}
