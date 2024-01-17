package shodan

import (
	"log"
	"os"
)

const ExceptionStatement string = "Missing search term and/or API key"

func (s *Client) KeyCheck() {
	// if len(os.Args) <= 2 {
	// 	log.Fatalln(exceptionStatement)
	// }

	if s.ApiKey == "" || s.ApiKey == "NOTDEFINED" {
		log.Fatalln(ExceptionStatement)
	}
}

func ArgCheck() {
	if len(os.Args) <= 2 {
		log.Fatalln(ExceptionStatement)
	}
}
