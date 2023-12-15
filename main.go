package main

import (
	"blackhat-go/scannertools"
	"fmt"
)

func main() {
	// scannertools.UserInputCheck()
	// scannertools.TestUserInput()
	userInput := scannertools.UserCommands()
	fmt.Println(userInput)
	// userInput := scannertools.ScannerUserInput()
	userHost := userInput["hostname"]
	userPorts := userInput["ports"]
	portsInt := scannertools.StringToIntPorts(&userPorts)
	scannertools.WorkerPoolScanTwoPorts(&userHost, &portsInt)
	// serverbind.BindServerPort()
}
