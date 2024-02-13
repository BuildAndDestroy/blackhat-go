package lib

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
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

type trivial struct {
}

func (t *trivial) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Executing trivial middleware")
	next(w, r)
}

func SimpleNegroni(r *mux.Router) {
	// n := negroni.Classic()
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(r)
	n.Use(&trivial{})
	n.Use(negroni.NewRecovery())
	http.ListenAndServe(":8000", n)
}

type badAuth struct {
	Username string
	Password string
}

func (b *badAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	if username != b.Username || password != b.Password {
		http.Error(w, "Unauthorized", 401)
		return
	}
	ctx := context.WithValue(r.Context(), "username", username)
	r = r.WithContext(ctx)
	next(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	fmt.Fprintf(w, "Hi %s\n", username)
}

func NegroniAuthExample(r *mux.Router) {
	r.HandleFunc("/hello", hello).Methods("GET")
	n := negroni.Classic()
	n.Use(&badAuth{
		Username: "admin",
		Password: "password",
	})
	n.UseHandler(r)
	http.ListenAndServe(":8000", n)
}
