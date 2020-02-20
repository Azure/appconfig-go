package appconfig_test

import (
	"testing"

	"github.com/Azure/appconfig-go"
)

func TestNew(t *testing.T) {

	cfg := appconfig.New("Endpoint=https://test.azconfig.io;Id=xZqB-l0-s0:uN7lw/eiMARw026xtjY7;Secret=RBWq9w8DZkcvDcsXUPrLi+2ua1zdPBkLtFq6sipHMSU=")
	if cfg.Endpoint != "https://test.azconfig.io" {
		t.Errorf("Expected %s, got %s", "https://test.azconfig.io", cfg.Endpoint)
	}

}
