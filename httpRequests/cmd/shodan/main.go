package main

import (
	"blackhat-go/httpRequests/shodan"
	"fmt"
)

func main() {
	shodan.TestGoogle()
	var baseUrl = shodan.BaseUrl
	fmt.Println(baseUrl)
}
