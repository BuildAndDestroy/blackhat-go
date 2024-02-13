package main

import (
	"blackhat-go/server/lib"
	"net/http"

	"github.com/gorilla/mux"
)

func TestFunctions() {
	// lib.SimpleServer()
	// lib.TestRouter()
	// lib.TestMiddleware()
	// lib.GetFoo()
	r := mux.NewRouter()
	// lib.GetFooHost(r)
	// lib.GetFoo(r)
	// lib.GetFooUser(r)
	lib.GetFooUserRegex(r)
	http.ListenAndServe(":8000", r)
}

func main() {
	TestFunctions()
}
