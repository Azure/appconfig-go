package appconfig

// ClientConfig contains optional configuration for creating
// a new AppConfig client
type ClientConfig struct {
	apiVersion string
}

// NewClientConfig creates a new client configuration, suitable
// for passing into appconfig.New()
func NewClientConfig(apiVersion string) *ClientConfig {
	return &ClientConfig{
		apiVersion: apiVersion,
	}
}

func (c *ClientConfig) APIVersion() string {
	if c == nil || c.apiVersion == "" {
		return "1.0"
	}
	return c.apiVersion
}
