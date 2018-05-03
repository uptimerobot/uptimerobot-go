package api

import (
	"os"
	"testing"
)

var monitorNameForTest = "testmonitor"
var monitorIDForTest = 0

func TestNewMonitor(t *testing.T) {
	c := makeClient(t)

	monitors := c.Monitors()

	var request = NewMonitorRequest{
		FriendlyName: monitorNameForTest,
		URL:          "http://www.apple.com",
		MonitorType:  HTTP,
	}
	response, err := monitors.New(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatalf("No monitor response: %v", response)
	}
	t.Logf("Monitor ID: %d", response.ID)
	t.Logf("New Monitor Response: %d", response.ID)
	monitorIDForTest = response.ID
}

func TestEditMonitor(t *testing.T) {
	c := makeClient(t)

	monitors := c.Monitors()

	var request = EditMonitorRequest{
		ID:           monitorIDForTest,
		FriendlyName: monitorNameForTest,
		URL:          "http://www.microsoft.com",
		MonitorType:  HTTP,
	}
	response, err := monitors.Edit(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatalf("No monitor response: %v", response)
	}
	t.Logf("Monitor ID: %d", response.ID)
	monitorIDForTest = response.ID
}

func TestDeleteMonitor(t *testing.T) {
	c := makeClient(t)

	monitors := c.Monitors()

	var request = DeleteMonitorRequest{
		ID: monitorIDForTest,
	}
	response, err := monitors.Delete(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatalf("No monitor response: %v", response)
	}
	t.Logf("Monitor ID: %d", response.ID)
	monitorIDForTest = response.ID
}

func TestGetMonitors(t *testing.T) {
	monitorID := os.Getenv("UPTIMEROBOT_MONITOR_ID")

	c := makeClient(t)

	monitors := c.Monitors()

	var request = GetMonitorsRequest{
		MonitorID:          monitorID,
		ResponseTimes:      1,
		ResponseTimesLimit: 1,
	}
	response, err := monitors.Get(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatalf("No monitor response: %v", response)
	}

	for _, monitor := range response.Monitors {
		t.Logf("Pagination Offset: %d", monitor.Pagination.Offset)
		t.Logf("Pagination Limit: %d", monitor.Pagination.Limit)
		t.Logf("Monitor ID: %d", monitor.ID)
		t.Logf("Monitor Friendly Name: %s", monitor.FriendlyName)
		t.Logf("Monitor URL: %s", monitor.URL)
		t.Logf("Monitor Status: %s", monitor.Status)
		t.Logf("Monitor Type: %s", monitor.Type)
		t.Logf("Monitor SubType: %s", monitor.SubType)
		t.Logf("Monitor Recent Response Time(msec): %d", monitor.ResponseTimes[0].Value)
	}
}
