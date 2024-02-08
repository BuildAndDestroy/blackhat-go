package main

import (
	"blackhat-go/metasploit-minimal/rpc"
	"fmt"
	"log"
	"os"
)

func ActiveSessions() {
	// Request active sessions in Metasploit
	var (
		host string = os.Getenv("MSFHOST")
		pass string = os.Getenv("MSFPASS")
		user string = "msf"
	)
	if host == "" || pass == "" {
		log.Fatalln("Missing required environment variable MSFHOST or MSFPASS")
	}

	msf, err := rpc.NewMetasploit(host, user, pass)

	if err != nil {
		log.Panicln(err)
	}

	defer msf.Logout()

	sessions, err := msf.SessionList()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Active Sessions:")
	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}

}

func main() {
	// Init the metasploit program
	ActiveSessions()
}
