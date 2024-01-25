package shodan

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type APIInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

func (s *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseUrl, s.ApiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Client) Credits() {
	info, err := s.APIInfo()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits: %d\nPlan: %s\n\n",
		info.QueryCredits,
		info.ScanCredits,
		info.Plan)
}

func (s *Client) HostIpPortSearch() {
	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
