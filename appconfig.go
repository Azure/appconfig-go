package appconfig

import (
	"strings"
)

// AppConfig contains the endpoint, ID and Secret for accessing
// Azure AppConfig service.  Create with New()
type AppConfig struct {
	Endpoint string
	ID       string
	Secret   string
}

// New creates a new AppConfig instance from the
// connection string provided in the Azure Portal.
func New(connectionString string) *AppConfig {

	b, i, s := parseConnection(connectionString)
	return &AppConfig{
		Endpoint: b,
		ID:       i,
		Secret:   s,
	}
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
