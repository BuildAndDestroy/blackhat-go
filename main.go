package main

import (
	"blackhat-go/scannertools"
	"blackhat-go/serverbind"
	"log"
	"os"
)

func main() {
	userInput := scannertools.UserCommands()
	// log.Println(userInput)
	switch userInput["command"] {
	case "Scanner":
		log.Println("[*] Initiating scanner")
		scannertools.InitScanner(userInput)
	case "Server":
		log.Println("[*] Initiating Server")
		serverbind.BindServerPort(userInput)
	case "Client":
		log.Println("[*] Initiating Client")
		log.Println("Client")
	case "Proxy":
		log.Println("[*] Initiating Proxy")
		serverbind.ProxyForward(userInput)
	default:
		log.Fatalln("Correct Subcommand does not exit")
		os.Exit(1)
	}
}
