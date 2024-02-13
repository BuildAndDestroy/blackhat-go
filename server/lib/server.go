package lib

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func HelloMiddleware(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello\n")
}

func SimpleServer() {
	http.HandleFunc("/hello", Hello)
	http.ListenAndServe(":8000", nil)
}

func TestRouter() {
	var r Router
	http.ListenAndServe(":8000", &r)
}

func TestMiddleware() {
	// Call the logger middleware while handling request.
	f := http.HandlerFunc(HelloMiddleware)
	l := logger{Inner: f}
	http.ListenAndServe(":8000", &l)
}
