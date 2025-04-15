package api

import (
	//"fmt"
	"testing"
)

func TestGetAccountDetails(t *testing.T) {
	c := makeClient(t)

	account := c.AccountDetails()

	details, err := account.Get()
	if err != nil {
		t.Fatal(err)
	}
	if details == nil {
		t.Fatalf("No account details: %v", details)
	}
	t.Logf("Monitor Limit  : %d", details.MonitorLimit)
	t.Logf("Up Monitors    : %d", details.UpMonitors)
	t.Logf("Down Monitors  : %d", details.DownMonitors)
	t.Logf("Paused Monitors: %d", details.PausedMonitors)
}
