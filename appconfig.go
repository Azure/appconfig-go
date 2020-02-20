package appconfig

import (
	"net/http"
	"strings"
)

// Client contains the endpoint, ID and Secret for accessing
// Azure Client service.  Create with New()
type Client struct {
	Endpoint string
	ID       string
	Secret   string
	client   *http.Client
	cfg      *ClientConfig
}

// New creates a new Client instance from the
// connection string provided in the Azure Portal.
func New(connectionString string, cfg *ClientConfig) *Client {

	b, i, s := parseConnection(connectionString)
	return &Client{
		Endpoint: b,
		ID:       i,
		Secret:   s,
		client:   http.DefaultClient,
		cfg:      cfg,
	}
}

// WithHTTPClient adds a pre-configured http.Client to the AppConfig
// struct.
func (a *Client) WithHTTPClient(c *http.Client) *Client {
	a.client = c
	return a
}

func parseConnection(cs string) (string, string, string) {

	var endpoint string
	var id string
	var secret string

	pieces := strings.Split(cs, ";")

	for _, piece := range pieces {

		var key string
		var value string
		eq := strings.Index(piece, "=")
		if eq > 0 {
			key = piece[0:eq]
			value = piece[eq+1:]
		}

		switch key {
		case "Endpoint":
			endpoint = value
		case "Secret":
			secret = value
		case "Id":
			id = value
		default:
		}

	}

	return endpoint, id, secret

}
