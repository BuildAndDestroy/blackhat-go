package main

import (
	"blackhat-go/scannertools"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// userInput := scannertools.UserCommands()
	// log.Println(userInput)
	scannertools.ArgLengthCheck()
	var command string = os.Args[1]
	scannertools.CommandCheck(command)
	var userCommand = flag.NewFlagSet(command, flag.ExitOnError)
	switch command {
	case "Scanner":
		var scanner scannertools.UserInputScanner
		scanner.SetFlagScanner(userCommand)
		userCommand.Parse(os.Args[2:])
		fmt.Println(scanner)
	case "Server":
		var server scannertools.UserInputServer
		server.SetFlagServer(userCommand)
		userCommand.Parse(os.Args[2:])
		fmt.Println(server)
	case "Client":
		var client scannertools.UserInputClient
		client.SetFlagClient(userCommand)
		userCommand.Parse(os.Args[2:])
		fmt.Println(client)
	case "Proxy":
		var proxy scannertools.UserInputProxy
		proxy.SetFlagProxy(userCommand)
		userCommand.Parse(os.Args[2:])
		fmt.Println(proxy)
	case "Netcat":
		var netcat scannertools.UserInputNetcat
		netcat.SetFlagNetcat(userCommand)
		userCommand.Parse(os.Args[2:])
		fmt.Println(netcat)
	default:
		log.Fatalln("Subcommand does not exist")
		os.Exit(1)
	}

	// switch userInput["command"] {
	// case "Scanner":
	// 	log.Println("[*] Initiating scanner")
	// 	scannertools.InitScanner(userInput)
	// case "Server":
	// 	log.Println("[*] Initiating Server")
	// 	serverbind.BindServerPort(userInput)
	// case "Client":
	// 	log.Println("[*] Initiating Client")
	// 	log.Println("Client")
	// case "Proxy":
	// 	log.Println("[*] Initiating Proxy")
	// 	serverbind.ProxyForward(userInput)
	// case "Netcat":
	// 	if userInput["bind"] == "true" { // Garbage, need to convert to a struct to handle strings, bools, etc.
	// 		log.Println("[*] Binding shell spawning for remote code execution")
	// 		serverbind.NcBind(userInput)
	// 	}
	// default:
	// 	log.Fatalln("Subcommand does not exist")
	// 	os.Exit(1)
	// }
}
