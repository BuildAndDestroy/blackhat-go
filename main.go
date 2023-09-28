package main

import (
	"blackhat-go/scannertools"
)

func main() {
	scannertools.UserInputCheck()
	userInput := scannertools.ScannerUserInput()
	userHost := userInput["hostname"]
	userPorts := userInput["ports"]
	portsInt := scannertools.StringToIntPorts(&userPorts)
	scannertools.WorkerPoolScanTwoPorts(&userHost, &portsInt)
}
