package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"

	"github.com/pablito-dev/family-calendar/utils"
)


func main()  {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/events"), getEvents)

	http.ListenAndServe("localhost:8080", mux)
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSON(w, []byte("{message: Hello world}"), 200)
}
