package ticket

import (
	"gopkg.in/mgo.v2/bson"
)

type Ticket struct {
	Id      bson.ObjectId `json:"id" bson:"id"`
	Event   bson.ObjectId `json:"event" bson:"event"`
	Scanned bool          `json:"scanned"`
}

/*
func main() {
	// MongoDB mLab connection

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"ds061938.mlab.com:61938"},
		Database: "cpts",
		Username: "test12",
		Password: "test12",
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
	r.HandleFunc("/api/ticket", createTicket).Methods("POST")
	r.HandleFunc("/api/ticket/{id}", updateTicket).Methods("PUT")
	r.HandleFunc("/api/ticket/{id}", deleteTicket).Methods("DELETE")
}

// Get all tickets
func (u *Ticket) getTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//TODO Get all Tickets from DB

	json.NewEncoder(w).Encode()
}

// Get ticket by param "id"
func (u *Ticket) getTicket(w http.ResponseWriter, r *http.Request) {
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
func (u *Ticket) createTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ticket Ticket

	_ = json.NewDecoder(r.Body).Decode(&ticket)

	ticket.Id = bson.NewObjectId()

	//TODO Add ticket to DB

}
*/