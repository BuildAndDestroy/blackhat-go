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
	proxy   string = "Proxy"
)

func UserCommands() map[string]string {
	// Let the user choose between client, server, scanner, or proxy.
	var (
		clientFlag         = flag.NewFlagSet(client, flag.ExitOnError)
		serverFlag         = flag.NewFlagSet(server, flag.ExitOnError)
		proxyFlag          = flag.NewFlagSet(proxy, flag.ExitOnError)
		serverArgPort      = serverFlag.String("port", "8000", "Port to bind to on this server/client.\nExample:\n    8000\n    1337")
		scannerFlag        = flag.NewFlagSet(scanner, flag.ExitOnError)
		scannerArgHostname = scannerFlag.String("host", "127.0.0.1", "Hostname or IP we want to scan")
		scannerArgPort     = scannerFlag.String("port", "0", "Port, or ports, to scan.\nExamples:\n    22\n    1-1000\n    22,443")
		proxyArgHost       = proxyFlag.String("target-host", "google.com", "Hostname to be our end target.")
		proxyArgTargetPort = proxyFlag.String("target-port", "80", "Port to query on our end target host.")
		proxyArgPort       = proxyFlag.String("port", "8000", "Port to bind to on this client.\nExample:\n    8000\n    1337")
	)

	if len(os.Args) <= 2 {
		fmt.Println("Expected 'Client', 'Server', 'Scanner', or 'Proxy' commands with a subcommand.")
		os.Exit(1)
	}

	var command string = os.Args[1]

	userInputMap := make(map[string]string)

	if command == client || command == server || command == scanner || command == proxy {
		switch command {

		case scanner:
			userInputMap["command"] = scanner
			scannerFlag.Parse(os.Args[2:])
			userInputMap["host"] = *scannerArgHostname
			userInputMap["ports"] = *scannerArgPort
			return userInputMap
		case client:
			userInputMap["command"] = client
			clientFlag.Parse(os.Args[2:])
			return userInputMap
		case server:
			userInputMap["command"] = server
			serverFlag.Parse(os.Args[2:])
			userInputMap["port"] = *serverArgPort
			return userInputMap
		case proxy:
			userInputMap["command"] = proxy
			proxyFlag.Parse(os.Args[2:])
			userInputMap["target-host"] = *proxyArgHost
			userInputMap["port"] = *proxyArgPort
			userInputMap["target-port"] = *proxyArgTargetPort
			return userInputMap
		default:
			fmt.Println("Missing subcommands")
			os.Exit(1)
		}
	} else {
		fmt.Println("Expected 'Client', 'Server', 'Scanner', or 'Proxy' commands with a subcommand.")
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
