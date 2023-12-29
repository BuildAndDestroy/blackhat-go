package main

import (
	"blackhat-go/scannertools"
	"blackhat-go/serverbind"
	"flag"
	"fmt"
	"log"
	"os"
)

func UserExecution() {
	var command string = os.Args[1]
	scannertools.CommandCheck(command)
	var userCommand = flag.NewFlagSet(command, flag.ExitOnError)

	switch command {
	case "Scanner":
		var scanner scannertools.UserInputScanner
		scanner.SetFlagScanner(userCommand)
		userCommand.Parse(os.Args[2:])
		scanner.InitScannerTwo()
	case "Server":
		var server serverbind.ServerBindUserInputServer
		server.SetFlagServer(userCommand)
		userCommand.Parse(os.Args[2:])
		server.BindServerPortTwo()
	case "Client":
		var client scannertools.UserInputClient
		client.SetFlagClient(userCommand)
		userCommand.Parse(os.Args[2:])
		fmt.Println(client)
	case "Proxy":
		var proxy serverbind.ServerBindUserInputProxy
		proxy.SetFlagProxy(userCommand)
		userCommand.Parse(os.Args[2:])
		proxy.ProxyForwardTwo()
	case "Netcat":
		var netcat serverbind.ServerBindUserInputNetcat
		netcat.SetFlagNetcat(userCommand)
		userCommand.Parse(os.Args[2:])
		if netcat.Bind {
			log.Println("[*] Binding shell spawning for remote code execution")
			netcat.NcBindTwo()
		}
		fmt.Println(netcat)
	default:
		log.Fatalln("Subcommand does not exist")
		os.Exit(1)
	}
}

// Originally main(), code cleanup resulted into structs, not maps.
// Leaving this here to see how this is done with maps.
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

func main() {
	scannertools.ArgLengthCheck()
	UserExecution()
}
