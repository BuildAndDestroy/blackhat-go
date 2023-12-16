package main

import (
	"blackhat-go/scannertools"
	"blackhat-go/serverbind"
	"log"
	"os"
)

func main() {
	userInput := scannertools.UserCommands()
	log.Println(userInput)
	switch userInput["command"] {
	case "Scanner":
		scannertools.InitScanner(userInput)
	case "Server":
		serverbind.BindServerPort(userInput)
	case "Client":
		log.Println("Client")
	default:
		log.Fatalln("Correct Subcommand does not exit")
		os.Exit(1)
	}
}
