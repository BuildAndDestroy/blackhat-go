package lib

import (
	"fmt"
	"net/http"
)

type Router struct {
	// router struct
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Constructor for router struct. Check for user input for
	// directories, else throw a 404.
	switch req.URL.Path {
	case "/a":
		fmt.Fprint(w, "Executing /a\n")
	case "/b":
		fmt.Fprint(w, "Executing /b\n")
	case "/c":
		fmt.Fprint(w, "Executing /c\n")
	default:
		http.Error(w, "404 Not Found", 404)
	}
}
