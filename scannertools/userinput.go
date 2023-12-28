package scannertools

import (
	"flag"
	"fmt"
	"os"
)

var EXCEPTIONSTATEMENT string = "Expected 'Client', 'Server', 'Scanner', 'Proxy', or 'Netcat' commands with a subcommand."

const (
	server  string = "Server"
	client  string = "Client"
	scanner string = "Scanner"
	echo    string = "Echo"
	proxy   string = "Proxy"
	netcat  string = "Netcat"
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

type UserInputServer struct {
	port int
}

func (uis *UserInputServer) SetFlagServer(fs *flag.FlagSet) {
	fs.IntVar(&uis.port, "port", 8000, "Port to bind to on this server/client.\nExample:\n    8000\n    1337")
}

type UserInputClient struct {
	host string
}

func (uic *UserInputClient) SetFlagClient(fs *flag.FlagSet) {
	fs.StringVar(&uic.host, "host", "127.0.0.1", "Hostname or IP we want to connect.")
}

type UserInputProxy struct {
	targetHost string
	targetPort int
	port       int
}

func (uip *UserInputProxy) SetFlagProxy(fs *flag.FlagSet) {
	fs.StringVar(&uip.targetHost, "target-host", "google.com", "Hostname to be our end target.")
	fs.IntVar(&uip.targetPort, "target-port", 80, "Port to query on our end target host.")
	fs.IntVar(&uip.port, "port", 8000, "Port to bind to on this client.\nExample:\n    8000\n    1337")
}

type UserInputScanner struct {
	host  string
	ports string
}

func (uis *UserInputScanner) SetFlagScanner(fs *flag.FlagSet) {
	fs.StringVar(&uis.host, "host", "127.0.0.1", "Hostname or IP we want to scan")
	fs.StringVar(&uis.ports, "port", "0", "Port, or ports, to scan.\nExamples:\n    22\n    1-1000\n    22,443")
}

type UserInputNetcat struct {
	bind bool
	port int
}

func (uin *UserInputNetcat) SetFlagNetcat(fs *flag.FlagSet) {
	fs.BoolVar(&uin.bind, "bind", false, "Create a bind shell. This will bind to the specified port, opening access to anyone who connects.")
	fs.IntVar(&uin.port, "port", 8000, "Bind to port on this host.")
}

func ArgLengthCheck() {
	if len(os.Args) == 1 {
		fmt.Println(EXCEPTIONSTATEMENT)
		os.Exit(1)
	}
}

// Check for user input matches our const, otherwise throw "exception" and exit
func CommandCheck(command string) {
	var exceptionStatement string = "Expected 'Client', 'Server', 'Scanner', 'Proxy', or 'Netcat' commands with a subcommand."
	if len(os.Args) <= 2 {
		fmt.Println(exceptionStatement)
		os.Exit(1)
	}

	if command == client || command == server || command == scanner || command == proxy || command == netcat {
		return
	} else {
		fmt.Println(exceptionStatement)
		os.Exit(1)
	}
}

func UserCommands() map[string]string {
	// Let the user choose between client, server, scanner, or proxy.
	var (
		clientFlag         = flag.NewFlagSet(client, flag.ExitOnError)
		serverFlag         = flag.NewFlagSet(server, flag.ExitOnError)
		proxyFlag          = flag.NewFlagSet(proxy, flag.ExitOnError)
		scannerFlag        = flag.NewFlagSet(scanner, flag.ExitOnError)
		netcatFlag         = flag.NewFlagSet(netcat, flag.ExitOnError)
		serverArgPort      = serverFlag.String("port", "8000", "Port to bind to on this server/client.\nExample:\n    8000\n    1337")
		scannerArgHostname = scannerFlag.String("host", "127.0.0.1", "Hostname or IP we want to scan")
		scannerArgPort     = scannerFlag.String("port", "0", "Port, or ports, to scan.\nExamples:\n    22\n    1-1000\n    22,443")
		proxyArgHost       = proxyFlag.String("target-host", "google.com", "Hostname to be our end target.")
		proxyArgTargetPort = proxyFlag.String("target-port", "80", "Port to query on our end target host.")
		proxyArgPort       = proxyFlag.String("port", "8000", "Port to bind to on this client.\nExample:\n    8000\n    1337")
		netcatBind         = netcatFlag.Bool("bind", false, "Create a bind shell. This will bind to the specified port, opening access to anyone who connects.")
		netcatArgPort      = netcatFlag.String("port", "8000", "Bind to port on this host.")
	)

	if len(os.Args) <= 2 {
		fmt.Println("Expected 'Client', 'Server', 'Scanner', 'Proxy', or 'Netcat' commands with a subcommand.")
		os.Exit(1)
	}

	var command string = os.Args[1]

	userInputMap := make(map[string]string)

	if command == client || command == server || command == scanner || command == proxy || command == netcat {
		switch command {

		case scanner:
			fmt.Println(userInputMap)
			userInputMap["command"] = scanner
			scannerFlag.Parse(os.Args[2:])
			userInputMap["host"] = *scannerArgHostname
			userInputMap["ports"] = *scannerArgPort
			fmt.Println(userInputMap)
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
		case netcat:
			userInputMap["command"] = netcat
			netcatFlag.Parse(os.Args[2:])
			userInputMap["port"] = *netcatArgPort
			if *netcatBind {
				userInputMap["bind"] = "true"
			}
			return userInputMap
		default:
			fmt.Println("Missing subcommands")
			os.Exit(1)
		}
	} else {
		fmt.Println("Expected 'Client', 'Server', 'Scanner', 'Proxy', or 'Netcat' commands with a subcommand.")
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
