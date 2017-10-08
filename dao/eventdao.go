package dao

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/pablito-dev/family-calendar/utils/databaseutils"
	"github.com/pablito-dev/family-calendar/utils/responseutils"
	"github.com/pablito-dev/family-calendar/models"
	"encoding/json"
)

func GetEvents(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var events []models.Event
		session := s.Copy()
		collection := databaseutils.SetCollectionInNewSession(session)
		defer session.Close()


		err := collection.Find(bson.M{}).All(&events)
		if err != nil {
			responseutils.ResponseWithError(w, err.Error(), 500)
			return
		}

		body, err := json.MarshalIndent(events, "", " ")
		if err != nil {
			responseutils.ResponseWithError(w, err.Error(), 500)
			return
		}

		responseutils.ResponseWithJSON(w, body, 200)
	}
}
