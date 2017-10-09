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

func RespondWithCode(w http.ResponseWriter, headers map[string]string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(code)
}