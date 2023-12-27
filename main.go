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
	case "Netcat":
		if userInput["bind"] == "true" { // Garbage, need to convert to a struct to handle strings, bools, etc.
			log.Println("[*] Binding shell spawning for remote code execution")
			serverbind.NcBind(userInput)
		}
	default:
		log.Fatalln("Subcommand does not exist")
		os.Exit(1)
	}
}
