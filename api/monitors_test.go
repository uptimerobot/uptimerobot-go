package api

import (
	"os"
	"strconv"
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
	envMonitorID := os.Getenv("UPTIMEROBOT_MONITOR_ID")

	if envMonitorID == "" {
		t.Skip("TestGetMonitors requires UPTIMEROBOT_MONITOR_ID env variable")
	}

	monitorID, err := strconv.Atoi(envMonitorID)
	if err != nil {
		t.Fatal(err)
	}

	c := makeClient(t)

	monitors := c.Monitors()

	var request = GetMonitorsRequest{
		MonitorID: monitorID,
	}
	response, err := monitors.Get(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatalf("No monitor response: %v", response)
	}

	monitor := response.Monitors[0]
	t.Logf("Monitor ID: %d", monitor.ID)
	t.Logf("Monitor Friendly Name: %s", monitor.FriendlyName)
	t.Logf("Monitor URL: %s", monitor.URL)
	t.Logf("Monitor Status: %s", monitor.Status)
	t.Logf("Monitor Type: %s", monitor.Type)
	t.Logf("Monitor SubType: %s", monitor.SubType)
	t.Logf("Monitor Recent Response Time(msec): %d", monitor.ResponseTimes[0].Value)
}
