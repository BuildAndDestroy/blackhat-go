package scannertools

import (
	"flag"
	"fmt"
	"os"
)

var EXCEPTIONSTATEMENT string = "Expected 'Client', 'Server', 'Scanner', 'Proxy', or 'Netcat' commands with a subcommand."

const (
	Server  string = "Server"
	Client  string = "Client"
	Scanner string = "Scanner"
	Echo    string = "Echo"
	Proxy   string = "Proxy"
	Netcat  string = "Netcat"
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
	Port int
}

func (uis *UserInputServer) SetFlagServer(fs *flag.FlagSet) {
	fs.IntVar(&uis.Port, "port", 8000, "Port to bind to on this server/client.\nExample:\n    8000\n    1337")
}

type UserInputClient struct {
	Host string
}

func (uic *UserInputClient) SetFlagClient(fs *flag.FlagSet) {
	fs.StringVar(&uic.Host, "host", "127.0.0.1", "Hostname or IP we want to connect.")
}

type UserInputProxy struct {
	TargetHost string
	TargetPort int
	Port       int
}

func (uip *UserInputProxy) SetFlagProxy(fs *flag.FlagSet) {
	fs.StringVar(&uip.TargetHost, "target-host", "google.com", "Hostname to be our end target.")
	fs.IntVar(&uip.TargetPort, "target-port", 80, "Port to query on our end target host.")
	fs.IntVar(&uip.Port, "port", 8000, "Port to bind to on this client.\nExample:\n    8000\n    1337")
}

type UserInputScanner struct {
	Host  string
	Ports string
}

func (uis *UserInputScanner) SetFlagScanner(fs *flag.FlagSet) {
	fs.StringVar(&uis.Host, "host", "127.0.0.1", "Hostname or IP we want to scan")
	fs.StringVar(&uis.Ports, "port", "0", "Port, or ports, to scan.\nExamples:\n    22\n    1-1000\n    22,443")
}

type UserInputNetcat struct {
	Bind    bool
	Reverse bool
	Port    int
	Address string
}

func (uin *UserInputNetcat) SetFlagNetcat(fs *flag.FlagSet) {
	fs.BoolVar(&uin.Bind, "bind", false, "Create a bind shell. This will bind to the specified port, opening access to anyone who connects.")
	fs.BoolVar(&uin.Reverse, "reverse", false, "Reverse shell. Provide Attacker IP or hostname to call back to.")
	fs.IntVar(&uin.Port, "port", 8000, "Bind to port on this host.")
	fs.StringVar(&uin.Address, "address", "127.0.0.1", "If bind shell, we listen on localhost, if reverse shell, add the attacker IP or host.")
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

	if command == Client || command == Server || command == Scanner || command == Proxy || command == Netcat {
		return
	} else {
		fmt.Println(exceptionStatement)
		os.Exit(1)
	}
}

func UserCommands() map[string]string {
	// Let the user choose between client, server, scanner, or proxy.
	var (
		clientFlag         = flag.NewFlagSet(Client, flag.ExitOnError)
		serverFlag         = flag.NewFlagSet(Server, flag.ExitOnError)
		proxyFlag          = flag.NewFlagSet(Proxy, flag.ExitOnError)
		scannerFlag        = flag.NewFlagSet(Scanner, flag.ExitOnError)
		netcatFlag         = flag.NewFlagSet(Netcat, flag.ExitOnError)
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

	if command == Client || command == Server || command == Scanner || command == Proxy || command == Netcat {
		switch command {

		case Scanner:
			fmt.Println(userInputMap)
			userInputMap["command"] = Scanner
			scannerFlag.Parse(os.Args[2:])
			userInputMap["host"] = *scannerArgHostname
			userInputMap["ports"] = *scannerArgPort
			fmt.Println(userInputMap)
			return userInputMap
		case Client:
			userInputMap["command"] = Client
			clientFlag.Parse(os.Args[2:])
			return userInputMap
		case Server:
			userInputMap["command"] = Server
			serverFlag.Parse(os.Args[2:])
			userInputMap["port"] = *serverArgPort
			return userInputMap
		case Proxy:
			userInputMap["command"] = Proxy
			proxyFlag.Parse(os.Args[2:])
			userInputMap["target-host"] = *proxyArgHost
			userInputMap["port"] = *proxyArgPort
			userInputMap["target-port"] = *proxyArgTargetPort
			return userInputMap
		case Netcat:
			userInputMap["command"] = Netcat
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
