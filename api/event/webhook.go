package event

import (
	"encoding/json"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Webhook struct {
	ID  bson.ObjectId `bson:"_id" json:"id"`
	URL string        `bson:"url" json:"url"`
}

type WebhookResponse struct {
	Content string         `json:"content"`
	Embeds  []WebhookEmbed `json:"embeds"`
}

type WebhookEmbed struct {
	Title string       `json:"title"`
	Image WebhookImage `json:"image"`
}

type WebhookImage struct {
	URL string `json:"url"`
}

func RegisterWebhook(w http.ResponseWriter, r *http.Request) (Webhook, error) {
	var wh Webhook

	// decoding post data
	decodeErr := json.NewDecoder(r.Body).Decode(&wh)
	if decodeErr != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return Webhook{}, decodeErr
	}

	wh.ID = bson.NewObjectId()

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
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return Webhook{}, dbErr
	}

	wh.ID = bson.NewObjectId()

	session.DB("cpts").C("webhooks").Insert(wh)

	return wh, nil
}
