package appconfig_test

import (
	"fmt"
	"testing"

	"github.com/Azure/appconfig-go"
)

func TestNew(t *testing.T) {

	client := appconfig.New("Endpoint=https://test.azconfig.io;Id=xZqB-l0-s0:uN7lw/eiMARw026xtjY7;Secret=RBWq9w8DZkcvDcsXUPrLi+2ua1zdPBkLtFq6sipHMSU=")
	if client.Endpoint != "https://test.azconfig.io" {
		t.Errorf("Expected %s, got %s", "https://test.azconfig.io", client.Endpoint)
	}

}

func TestUsage(t *testing.T) {
	client := appconfig.New("Endpoint=https://test.azconfig.io;Id=xZqB-l0-s0:uN7lw/eiMARw026xtjY7;Secret=RBWq9w8DZkcvDcsXUPrLi+2ua1zdPBkLtFq6sipHMSU=")
	keys := client.GetAllKeys()
	fmt.Println(keys)
}
