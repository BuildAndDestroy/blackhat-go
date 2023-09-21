package main

import (
	"blackhat-go/scannertools"
	"fmt"
)

func main() {
	// host := "scanme.nmap.org"
	// host := "127.0.0.1"
	// go scannertools.SinglePort(host, 80)
	// scannertools.TenTwentyFourPorts(host)
	// scannertools.WorkerPoolScanTwo(host)
	scannertools.UserInputCheck()
	// scannertools.TestUserInput()
	gotheem := scannertools.ScannerUserInput()
	portsInt := scannertools.ConvertArrayPortsToInt(gotheem["ports"])
	fmt.Println(portsInt)
}
