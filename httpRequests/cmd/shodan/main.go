package main

import (
	"blackhat-go/httpRequests/shodan"
	"flag"
	"os"
)

func main() {
	// shodan.TestGoogle()
	shodan.ArgCheck()
	var searchTerm string = os.Args[1]
	var userInputSearchTermFlag = flag.NewFlagSet(searchTerm, flag.ExitOnError)

	var shodanClient shodan.Client
	shodanClient.SetFlagShodanKey(userInputSearchTermFlag)
	userInputSearchTermFlag.Parse(os.Args[2:])
	shodanClient.KeyCheck()
	s := shodan.New(shodanClient.ApiKey)
	s.Credits()
	s.HostIpPortSearch()
}
