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
		friendlyName: monitorNameForTest,
		url:          "http://www.apple.com",
		monitorType:  Http,
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
		id:           monitorIdForTest,
		friendlyName: monitorNameForTest,
		url:          "http://www.microsoft.com",
		monitorType:  Http,
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
		id: monitorIdForTest,
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
