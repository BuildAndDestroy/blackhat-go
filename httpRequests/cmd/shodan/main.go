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
	// log.Println(shodanClient.ApiKey)

	// apiKey := os.Getenv("SHODAN_API_KEY")
	// fmt.Println(apiKey)
	s := shodan.New(shodanClient.ApiKey)
	s.Credits()
	s.HostIpPortSearch()
	// hostSearch, err := s.HostSearch(os.Args[1])
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// for _, host := range hostSearch.Matches {
	// 	fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	// }
}
