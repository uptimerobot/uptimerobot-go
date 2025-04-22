package api

import (
	"net/http"
	"os"
	"testing"
)

func makeClient(t *testing.T) *Client {
	// Create client
	client, err := NewClient(os.Getenv("UPTIMEROBOT_KEY"), http.DefaultClient)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	return client
}
