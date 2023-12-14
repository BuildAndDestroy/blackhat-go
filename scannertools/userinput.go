package scannertools

import (
	"flag"
	"fmt"
	"os"
)

func TestUserInput() {
	// Good read for covering user input:
	// https://gobyexample.com/command-line-flags

	wordPtr := flag.String("word", "foo", "User defined string")
	numbPtr := flag.Int("number", 42, "An integer")
	forkPtr := flag.Bool("bool", false, "A bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "A string variable")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("bool:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("args:", flag.Args())
}

const (
	server  string = "Server"
	client  string = "Client"
	scanner string = "Scanner"
	echo    string = "Echo"
)

func UserCommands() map[string]string {
	// Let the user choose between client, server, or scanner.
	var (
		clientFlag         = flag.NewFlagSet(client, flag.ExitOnError)
		serverFlag         = flag.NewFlagSet(server, flag.ExitOnError)
		scannerFlag        = flag.NewFlagSet(scanner, flag.ExitOnError)
		scannerArgHostname = scannerFlag.String("Hostname", "127.0.0.1", "Hostname or IP we want to scan")
		scannerArgPort     = scannerFlag.String("Port", "0", "Port, or ports, to scan.\nExamples:\n    22\n    1-1000\n    22,443")
	)

	if len(os.Args) <= 2 {
		fmt.Println("Expected 'Client', 'Server', or 'Scanner' commands with a subcommand.")
		os.Exit(1)
	}

	var command string = os.Args[1]

	userInputMap := make(map[string]string)

	if command == client || command == server || command == scanner {
		switch command {

		case scanner:
			scannerFlag.Parse(os.Args[2:])
			userInputMap["hostname"] = *scannerArgHostname
			userInputMap["ports"] = *scannerArgPort
			return userInputMap
		case client:
			clientFlag.Parse(os.Args[2:])
		case server:
			serverFlag.Parse(os.Args[2:])
		default:
			fmt.Println("Missing subcommands")
			os.Exit(1)
		}
	} else {
		fmt.Println("Expected 'Client', 'Server', or 'Scanner' commands with a subcommand.")
		os.Exit(1)
	}
	return userInputMap
}

func ScannerUserInput() map[string]string {

	hostPtr := flag.String("Hostname", "127.0.0.1", "Hostname or IP we want to scan")
	portPtr := flag.String("Port", "0", "Port, or ports, to scan.\nExamples:\n    22\n    1-1000\n    22,443")
	flag.Parse()

	userInputMap := make(map[string]string)
	userInputMap["hostname"] = *hostPtr
	userInputMap["ports"] = *portPtr

	return userInputMap
}

func UserInputCheck() {
	if len(os.Args) == 1 {
		fmt.Println("No arguments provided.")
		os.Exit(1)
	}
}
