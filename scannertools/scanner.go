package scannertools

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

func SinglePort(host string, port int) {
	fmt.Printf("[*] Single port scan for port %d\n", port)
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err == nil {
		message := fmt.Sprintf("Connection to %d successful!", port)
		fmt.Println(message)
	}
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
}

func TenTwentyFourPorts(host string) {
	fmt.Printf("[*] Scanning the first 1024 ports\n")
	for i := 0; i <= 1024; i++ {
		address := fmt.Sprintf("%s:%d", host, i)
		// fmt.Println(address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// Port is filtered or closed
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}

func SyncronizeScanGroups(host string) {
	// This is incorrect, you will see inconsistent results.
	fmt.Println("[*] Using SyncGroup to scan")
	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}

func Worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func WorkerPoolScan() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go Worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}

func WorkerTwo(host string, ports, results chan int) {
	// From the book, a worker function that scans a port
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func WorkerPoolScanTwo(host string) {
	// From the book, worker pool that scans ports 1 to 1024 via channels.
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go WorkerTwo(host, ports, results)
	}
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

func WorkerPoolScanTwoPorts(host *string, userPorts *[]int) {
	// From the book, just modified to accept user input of ports
	// 100 buffer that will run buckets of ports.
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go WorkerTwo(*host, ports, results)
	}
	go func() {
		for _, i := range *userPorts {
			ports <- i
		}
	}()

	for range *userPorts {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

func InitScanner(mappedUserInput map[string]string) {
	// Initialize the scanner, begin scanning
	userHost := mappedUserInput["host"]
	userPorts := mappedUserInput["ports"]
	portsInt := StringToIntPorts(&userPorts)
	WorkerPoolScanTwoPorts(&userHost, &portsInt)
}

func (uis *UserInputScanner) InitScannerTwo() {
	portsInt := StringToIntPorts(&uis.Ports)
	WorkerPoolScanTwoPorts(&uis.Host, &portsInt)
}
