package iplocation

import (
	"context"
	"testing"
)

// TestClient_GetIPAddress is the simplest example.
func TestClient_GetIPAddress(t *testing.T) {
	locator, err := IPAddressLocator(context.Background())
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", locator)
}

// TestClient_IPLocatorByIP is the simplest example.
func TestClient_IPLocatorByIP(t *testing.T) {
	locator, err := IPLocatorByIP(context.Background(), "8.8.8.8")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", locator)
}
