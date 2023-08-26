package main

import (
	"scannertools/scannertools"
)

func main() {
	// host := "scanme.nmap.org"
	host := "127.0.0.1"
	// go scannertools.SinglePort(host, 80)
	// scannertools.TenTwentyFourPorts(host)
	scannertools.WorkerPoolScanTwo(host)
}
