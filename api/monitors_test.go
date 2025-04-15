package api

import (
	"os"
	"strconv"
	"testing"
)

var monitorNameForTest = "testmonitor"
var monitorIdForTest = 0

func TestNewMonitor(t *testing.T) {
	c := makeClient(t)

	monitors := c.Monitors()

	var request = NewMonitorRequest{
		FriendlyName: monitorNameForTest,
		Url:          "http://www.apple.com",
		MonitorType:  Http,
	}
	response, err := monitors.New(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatalf("No monitor response: %v", response)
	}
	t.Logf("Monitor ID: %d", response.ID)
	monitorIdForTest = response.ID
}

func TestEditMonitor(t *testing.T) {
	c := makeClient(t)

	monitors := c.Monitors()

	var request = EditMonitorRequest{
		Id:           monitorIdForTest,
		FriendlyName: monitorNameForTest,
		Url:          "http://www.microsoft.com",
	}
	response, err := monitors.Edit(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatalf("No monitor response: %v", response)
	}
	t.Logf("Monitor ID: %d", response.ID)
	monitorIdForTest = response.ID
}

func TestDeleteMonitor(t *testing.T) {
	c := makeClient(t)

	monitors := c.Monitors()

	var request = DeleteMonitorRequest{
		Id: monitorIdForTest,
	}
	response, err := monitors.Delete(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatalf("No monitor response: %v", response)
	}
	t.Logf("Monitor ID: %d", response.ID)
	monitorIdForTest = response.ID
}

func TestGetMonitors(t *testing.T) {
	envMonitorId := os.Getenv("UPTIMEROBOT_MONITOR_ID")

	if envMonitorId == "" {
		t.Skip("TestGetMonitors requires UPTIMEROBOT_MONITOR_ID env variable")
	}

	monitorId, err := strconv.Atoi(envMonitorId)
	if err != nil {
		t.Fatal(err)
	}

	c := makeClient(t)

	monitors := c.Monitors()

	var request = GetMonitorsRequest{
		MonitorId: monitorId,
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
	t.Logf("Monitor Recent Response Time(msec): %d", monitor.ResponseTimeList.ResponseTimes[0].Value)
}
