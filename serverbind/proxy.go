package serverbind

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handle(src net.Conn, host string, port string) {
	// Connect to user requested host
	var address string = fmt.Sprintf("%s:%s", host, port)
	dst, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Unable to reach %s\n", address)
	}
	defer dst.Close()

	// Run in goroutine to avoid io.Copy from blocking
	go func() {
		// Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func ProxyForward(mappedUserInput map[string]string) {
	// Drop on compromised host.
	// Listen on local port.
	// Connect to remote host. Allow for client to proxy through me to another host.
	var (
		targetHost string = mappedUserInput["target-host"]
		port       string = mappedUserInput["port"]
		targetPort string = mappedUserInput["target-port"]
		address    string = fmt.Sprintf(":%s", port)
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
		go handle(conn, targetHost, targetPort)
	}
}
