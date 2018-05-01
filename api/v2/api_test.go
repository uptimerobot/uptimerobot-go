package v2

import (
	"os"
	"testing"
)

func makeClient(t *testing.T) *Client {
	// Create client
	client, err := NewClient(os.Getenv("UPTIMEROBOT_API_KEY"))
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	return client
}
