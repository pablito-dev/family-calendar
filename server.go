package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello world", html.EscapeString(request.URL.Path))
}

func main()  {
	http.HandleFunc( "/", handleIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
