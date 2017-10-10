package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"

	"github.com/pablito-dev/family-calendar/utils/databaseutils"
	"github.com/pablito-dev/family-calendar/dao"
)


func main()  {
	session := databaseutils.ConnectToDatabase()
	databaseutils.PrepareDatabase(session)
	defer session.Close()

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/events"), dao.GetEvents(session))
	mux.HandleFunc(pat.Post("/events"), dao.CreateEvent(session))
	mux.HandleFunc(pat.Get("/events/:eventId"), dao.GetEventById(session))
	mux.HandleFunc(pat.Delete("/events/:eventId"), dao.DeleteEvent(session))
	mux.HandleFunc(pat.Put("/events/:eventId"), dao.UpdateEvent(session))

	http.ListenAndServe("localhost:8080", mux)
}
