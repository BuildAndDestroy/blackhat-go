package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Start TCP server on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server started. Waiting for connections...")

	// Accept incoming connections
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected. Receiving file...")

	// Create a new file to write the received content
	outputFile, err := os.Create("received.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	// Create a buffer reader for the TCP connection
	reader := bufio.NewReader(conn)
	// fmt.Println(reader)

	// Create a buffer writer for writing to the file
	writer := bufio.NewWriter(outputFile)

	// Read data from the TCP connection and write to the file
	_, err = io.Copy(writer, reader)
	if err != nil {
		fmt.Println("Error receiving file:", err)
		return
	}

	// Flush any unwritten data from the writer to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing data to file:", err)
		return
	}

	fmt.Println("File received successfully!")
}

// Todo: Change filename as whatever the user wants
