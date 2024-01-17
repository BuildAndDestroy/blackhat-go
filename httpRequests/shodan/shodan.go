package shodan

import "flag"

const BaseUrl = "https://api.shodan.io"

type Client struct {
	ApiKey string
}

func New(apiKey string) *Client {
	return &Client{ApiKey: apiKey}
}

func (a *Client) SetFlagShodanKey(fs *flag.FlagSet) {
	fs.StringVar(&a.ApiKey, "shodan-api-key", "NOTDEFINED", "Your shodan API key. This flag as plaintext or define as a global variable.\n    Example: --shodan-api-key $SHODAN_API_KEY")
}
