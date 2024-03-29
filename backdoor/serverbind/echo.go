package serverbind

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"

	"github.com/BuildAndDestroy/blackhat-go/backdoor/scannertools"
)

func copyEcho(conn net.Conn) {
	// Copy data from io.Reader to io.Writer via io.Copy
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data.")
	}
}

func bufioEcho(conn net.Conn) {
	// buffEcho is a handler function that simply echos received data.
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data.")
	}
	log.Printf("Read %d bytes: %s", len(s), s)

	log.Printf("Writing data.")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
}

func echo(conn net.Conn) {
	// echo is a handler function that simply echos received data.
	defer conn.Close()

	b := make([]byte, 512)
	for {
		// Receive data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected.")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		// Send data via conn.Write.
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func BindServerPort(mappedUserInput map[string]string) {
	// Bind to user specified TCP port on all interfaces.
	var (
		mappedPort   string = mappedUserInput["port"]
		stringPort   string = mappedPort
		listenerPort string = ":" + stringPort
	)

	listener, err := net.Listen("tcp", listenerPort)
	if err != nil {
		log.Fatalln("Unable to bind to port " + stringPort)
	}
	log.Println("Listening on port " + stringPort)
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Printf("Received connection from %s!\n", conn.RemoteAddr().String())
		if err != nil {
			log.Fatalln("Unable to accept connection.")
		}
		// Handle the connection. Using goroutine for concurrency.
		// go echo(conn)
		// go bufioEcho(conn)
		go copyEcho(conn) // Much more stable
	}
}

type ServerBindUserInputServer struct {
	scannertools.UserInputServer
}

func (uis *ServerBindUserInputServer) BindServerPortTwo() {
	var (
		mappedPort   string = strconv.Itoa(uis.Port)
		stringPort   string = mappedPort
		listenerPort string = ":" + stringPort
	)

	listener, err := net.Listen("tcp", listenerPort)
	if err != nil {
		log.Fatalln("Unable to bind to port " + stringPort)
	}
	log.Println("Listening on port " + stringPort)
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Printf("Received connection from %s!\n", conn.RemoteAddr().String())
		if err != nil {
			log.Fatalln("Unable to accept connection.")
		}
		// Handle the connection. Using goroutine for concurrency.
		// go echo(conn)
		// go bufioEcho(conn)
		go copyEcho(conn) // Much more stable
	}
}
