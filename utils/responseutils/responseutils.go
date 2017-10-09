package responseutils

import (
	"net/http"
	"fmt"
)

func RespondWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func RespondWithError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprint(w, "{message: %q", message)
}

func RespondWithLocationHeader(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Location", r.URL.Path + "/" + id)
	w.WriteHeader(http.StatusCreated)
}