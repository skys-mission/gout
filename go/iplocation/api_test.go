package iplocation

import (
	"context"
	"testing"
)

func newTestClient() Client {
	return New()
}

func TestClient_GetIPAddress(t *testing.T) {
	testClient := newTestClient()
	locator, err := testClient.GetIPAddress(context.Background())
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", locator)
}

func TestClient_IPLocator(t *testing.T) {
	testClient := newTestClient()
	locator, err := testClient.IPLocator(context.Background(), "8.8.8.8")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("%+v", locator)
}
