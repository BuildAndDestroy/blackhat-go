package main

import (
	"net/http"

	"github.com/BuildAndDestroy/blackhat-go/server/lib"

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
	// lib.GetFooUserRegex(r)
	http.ListenAndServe(":8000", r)
}

func TestNegroniFunctions() {
	r := mux.NewRouter()
	// lib.SimpleNegroni(r)
	lib.NegroniAuthExample(r)
}

func main() {
	// TestFunctions()
	// TestNegroniFunctions()
	lib.CallHelloHTMLPage()
}
