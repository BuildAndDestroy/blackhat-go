package serverbind

import (
	"blackhat-go/scannertools"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

func handle(src net.Conn, host string, port int) {
	// Connect to user requested host
	var address string = fmt.Sprintf("%s:%d", host, port)
	dst, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Unable to reach %s\n", address)
	}
	defer dst.Close()

	// Run in goroutine to avoid io.Copy from blocking
	go func() {
		// Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Println(err) // Book says to run log.Fatalln, changing to Print since this should not kill our proxy.
		}
	}()
	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		log.Println(err) // Book says to run log.Fatalln, changing to Print since this should not kill our proxy.
	}
}

func Atoi(s string) int {
	// Convert string to int, strip the error.
	value, _ := strconv.Atoi(s)
	return value
}

func ProxyForward(mappedUserInput map[string]string) {
	// Drop on compromised host.
	// Listen on local port.
	// Connect to remote host. Allow for client to proxy through me to another host.
	var (
		targetHost    string = mappedUserInput["target-host"]
		port          string = mappedUserInput["port"]
		targetPort    string = mappedUserInput["target-port"] // This returns (string, error) example: 443 <nil>
		targetPortInt int    = Atoi(targetPort)
		address       string = fmt.Sprintf(":%s", port)
	)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Unable to bind to port %s", address)
	}

	for {
		conn, err := listener.Accept()
		log.Printf("Received connection from %s!\n", conn.RemoteAddr().String())
		if err != nil {
			log.Fatalln("Unable to accept connection.")
		}
		go handle(conn, targetHost, targetPortInt)
	}
}

type ServerBindUserInputProxy struct {
	scannertools.UserInputProxy
}

func (uip *ServerBindUserInputProxy) ProxyForwardTwo() {
	// Drop on compromised host.
	// Listen on local port.
	// Connect to remote host. Allow for client to proxy through me to another host.
	var (
		targetHost string = uip.TargetHost
		port       int    = uip.Port
		targetPort int    = uip.TargetPort
		address    string = fmt.Sprintf(":%d", port)
	)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Unable to bind to port %s", address)
	} else {
		log.Printf("Listening on port %d", port)
	}

	for {
		conn, err := listener.Accept()
		log.Printf("Received connection from %s!\n", conn.RemoteAddr().String())
		if err != nil {
			log.Fatalln("Unable to accept connection.")
		} else {
			log.Printf("Reaching out to %s on port %d", targetHost, targetPort)
		}
		go handle(conn, targetHost, targetPort)
	}
}
