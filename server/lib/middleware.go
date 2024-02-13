package lib

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type logger struct {
	// Logger struct for middleware
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Logger constructor
	log.Println("start")
	l.Inner.ServeHTTP(w, r)
	log.Println("finish")
}

func GetFoo(r *mux.Router) {
	// Run a GET request to /foo, print Hi foo
	// r := mux.NewRouter()
	r.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hi foo\n")
	}).Methods("GET")
}

func GetFooHost(r *mux.Router) {
	// r := mux.NewRouter()
	r.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hi foo\n")
	}).Methods("GET").Host("www.foo.com")
}

func GetFooUser(r *mux.Router) {
	// Client requests a user, say hi to user
	r.HandleFunc("/user/{user}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "Hi %s\n", user)
	}).Methods("GET")
}

func GetFooUserRegex(r *mux.Router) {
	r.HandleFunc("/user/{user:[a-z]+}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "Hi %s\n", user)
	}).Methods("GET")
}
