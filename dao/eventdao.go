package dao

import (
	"net/http"
	"encoding/json"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goji.io/pat"
	"github.com/satori/go.uuid"

	"github.com/pablito-dev/family-calendar/utils/databaseutils"
	"github.com/pablito-dev/family-calendar/utils/responseutils"
	"github.com/pablito-dev/family-calendar/models"
)

func GetEvents(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var events []models.Event
		session := s.Copy()
		collection := databaseutils.SetCollectionInNewSession(session)
		defer session.Close()


		err := collection.Find(bson.M{}).All(&events)
		if err != nil {
			responseutils.RespondWithError(w, err.Error(), 500)
			return
		}

		body, err := json.MarshalIndent(events, "", " ")
		if err != nil {
			responseutils.RespondWithError(w, err.Error(), 500)
			return
		}

		responseutils.RespondWithJSON(w, body, 200)
	}
}

func GetEventById(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var event models.Event

		id := pat.Param(r, "eventId")
		session := s.Copy()
		collection := databaseutils.SetCollectionInNewSession(session)
		defer session.Close()

		err := collection.Find(bson.M{"id": id}).One(&event)
		if err != nil {
			responseutils.RespondWithError(w, err.Error(), 500)
			return
		}

		body, err := json.MarshalIndent(event, "", " ")
		if err != nil {
			responseutils.RespondWithError(w, err.Error(), 500)
			return
		}

		responseutils.RespondWithJSON(w, body, 200)
	}
}

func CreateEvent(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newEvent models.Event
		session := s.Copy()
		collection := databaseutils.SetCollectionInNewSession(session)
		defer session.Close()

		err := json.NewDecoder(r.Body).Decode(&newEvent)
		if err != nil {
			responseutils.RespondWithError(w, err.Error(), 500)
		}
		newEvent.Id = uuid.NewV4().String()

		err = collection.Insert(newEvent)
		if err != nil {
			responseutils.RespondWithError(w, err.Error(), 500)
		}

		var headers = make(map[string]string)

		headers["Location"] = r.URL.Path + "/" + newEvent.Id
		responseutils.RespondWithCode(w, headers, 201)
	}
}

func DeleteEvent (s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		collection := databaseutils.SetCollectionInNewSession(session)
		id := pat.Param(r, "eventId")

		defer session.Close()

		err := collection.Remove(bson.M{"id": id})
		if err != nil {
			responseutils.RespondWithError(w, err.Error(), 500)
		}

		responseutils.RespondWithCode(w, nil, 204)
	}
}