package api

import (
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
		t.Fatal("No monitor response: %v", response)
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
		MonitorType:  Http,
	}
	response, err := monitors.Edit(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatal("No monitor response: %v", response)
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
		t.Fatal("No monitor response: %v", response)
	}
	t.Logf("Monitor ID: %d", response.ID)
	monitorIdForTest = response.ID
}
