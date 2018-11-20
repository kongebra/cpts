package test

import (
	"fmt"
	"github.com/kongebra/cpts/api/db"
	"github.com/kongebra/cpts/api/event"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"testing"
)

func TestEventDatabaseInit(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	index := mgo.Index{
		Key: []string{"name"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}

	err = session.DB(db.DATABASE).C(event.COLLECTION).EnsureIndex(index)

	if err != nil {
		t.Error("Could not ensure index")
	}
}

func TestEventDatabaseInsert(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	e := event.Event{
		Name: "HusetLan v2019",
		Description: "Some description",
		Date: event.TimeInterval{
			Start: "2019-03-11T16:00:00Z",
			End: "2019-03-13T12:00:00Z",
		},
	}

	err = session.DB(db.DATABASE).C(event.COLLECTION).Insert(e)

	if err != nil {
		t.Error("Could not insert data")
	}
}

func TestEventDatabaseGetSingle(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	var e event.Event

	err = session.DB(db.DATABASE).C(event.COLLECTION).Find(bson.M{"name": "HusetLan v2019"}).One(&e)

	if err != nil {
		t.Error("Could not get single document from database")
	}

	fmt.Println(e)
}

func TestEventDatabaseGetAll(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	var e []event.Event

	err = session.DB(db.DATABASE).C(event.COLLECTION).Find(bson.M{}).All(&e)

	if err != nil {
		t.Error("Could not get single document from database")
	}

	fmt.Println(e)
}

func TestEventDatabaseCount(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	count, err := session.DB(db.DATABASE).C(event.COLLECTION).Count()

	if err != nil {
		t.Error("Could not count the amount of documents in the database")
	}

	if count <= 0 {
		t.Error("Should be more than 0 documents in the database")
	}

	log.Printf("Document count: %d", count)
}

func TestEventDatabaseRemove(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	id := bson.NewObjectId()

	e := event.Event{
		Id: id,
		Name: "HusetLan h2019",
		Description: "Some description",
		Date: event.TimeInterval{
			Start: "2019-10-11T16:00:00Z",
			End: "2019-10-13T12:00:00Z",
		},
	}

	err = session.DB(db.DATABASE).C(event.COLLECTION).Insert(e)

	if err != nil {
		t.Error("Could not insert data")
	}

	count0, err := session.DB(db.DATABASE).C(event.COLLECTION).Count()

	if err != nil {
		t.Error("Could not count collection")
	}

	err = session.DB(db.DATABASE).C(event.COLLECTION).Remove(bson.M{"_id": id})

	if err != nil {
		t.Error("Could not remove for collection")
	}

	count1, err := session.DB(db.DATABASE).C(event.COLLECTION).Count()

	if err != nil {
		t.Error("Could not count collection")
	}

	if count0 == count1 {
		t.Error("First and second count is equal, second should be one less")
	}
}

func TestEventDatabaseAddParticipant(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	var e event.Event

	err = session.DB(db.DATABASE).C(event.COLLECTION).Find(bson.M{"name": "HusetLan v2019"}).One(&e)

	if err != nil {
		t.Error("Could not find one from the collection")
	}

	e.Participants = append(e.Participants, bson.NewObjectId())
	e.Participants = append(e.Participants, bson.NewObjectId())
	e.Participants = append(e.Participants, bson.NewObjectId())

	err = session.DB(db.DATABASE).C(event.COLLECTION).Update(bson.M{"_id": e.Id}, e)

	if err != nil {
		t.Error("Could not update document")
	}
}

func TestEventDatabaseRemoveParticipant(t *testing.T) {
	session, err := mgo.Dial(db.URL)

	defer session.Close()

	if err != nil {
		t.Error("Could not dial to database")
	}

	var e event.Event

	err = session.DB(db.DATABASE).C(event.COLLECTION).Find(bson.M{"name": "HusetLan v2019"}).One(&e)

	if err != nil {
		t.Error("Could not find one from the collection")
	}

	before := len(e.Participants)

	if err != nil {
		t.Error("Could not count collection")
	}

	var p = e.Participants[0].Hex()

	err = session.DB(db.DATABASE).C(event.COLLECTION).Update(bson.M{"_id": e.Id}, bson.M{"$pull": bson.M{"participants": bson.ObjectIdHex(p)}})

	if err != nil {
		t.Error("Could not remove document from collection based on participant")
	}

	err = session.DB(db.DATABASE).C(event.COLLECTION).Find(bson.M{"name": "HusetLan v2019"}).One(&e)

	if err != nil {
		t.Error("Could not find one from the collection")
	}

	after := len(e.Participants)

	if before <= after {
		t.Error("After count should be smaller than the count before the pull")
	}
}

