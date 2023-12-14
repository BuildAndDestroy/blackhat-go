package serverbind

import (
	"io"
	"log"
	"net"
	"strconv"
)

// echo is a handler function that simply echos received data.
func echo(conn net.Conn) {
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

func BindServerPort() {
	// Bind to TCP port 20080 on all interfaces.
	var (
		port         int    = 20080
		stringPort   string = strconv.Itoa(port)
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
		log.Println("Recieved connection!")
		if err != nil {
			log.Fatalln("Unable to accept connection.")
		}
		// Handle the connection. Using goroutine for concurrency.
		go echo(conn)
	}
}
