package shodan

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func TestGoogle() {
	// Simple GET request against google
	resp, err := http.Get("https://google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}
