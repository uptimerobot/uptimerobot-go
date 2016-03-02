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
		t.Fatal("No monitor response: %v", response)
	}
	t.Logf("Monitor ID: %d", response.ID)
	monitorIdForTest = response.ID
}

func TestNewMonitorWithAlertContacts(t *testing.T) {
	c := makeClient(t)

	alertContacts := c.AlertContacts()

	var alertRequest = NewAlertContactRequest{
		AlertContactType:  AlertContactEmail,
		AlertContactValue: "monitorcontacttest@test.com",
	}
	alertResponse, err := alertContacts.New(alertRequest)
	if err != nil {
		t.Fatal(err)
	}
	alertID := alertResponse.ID

	monitors := c.Monitors()

	var monitorRequest = NewMonitorRequest{
		FriendlyName:    "monitorwithalerts",
		Url:             "http://www.google.com",
		MonitorType:     Http,
		AlertContactIDs: []int{alertID},
	}
	monitorResponse, err := monitors.New(monitorRequest)
	if err != nil {
		t.Fatal(err)
	}
	if monitorResponse == nil {
		t.Fatal("No monitor response: %v", monitorResponse)
	}
	monitorID := monitorResponse.ID

	_ = monitorID
	deleteMonitorRequest := DeleteMonitorRequest{
		Id: monitorID,
	}
	_, err = monitors.Delete(deleteMonitorRequest)
	if err != nil {
		t.Fatal(err)
	}

	deleteAlertRequest := DeleteAlertContactRequest{
		ID: alertID,
	}
	_, err = alertContacts.Delete(deleteAlertRequest)
	if err != nil {
		t.Fatal(err)
	}

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

func TestGetMonitors(t *testing.T) {
	envMonitorId := os.Getenv("UPTIMEROBOT_MONITOR_ID")

	if envMonitorId == "" && monitorIdForTest == 0 {
		t.Skip("TestGetMonitors requires UPTIMEROBOT_MONITOR_ID env variable")
	}
	var monitorId int
	var err error

	if envMonitorId != "" {
		monitorId, err = strconv.Atoi(envMonitorId)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		monitorId = monitorIdForTest
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
		t.Fatal("No monitor response: %v", response)
	}

	monitor := response.Monitors[0]
	t.Logf("Monitor ID: %d", monitor.ID)
	t.Logf("Monitor Friendly Name: %s", monitor.FriendlyName)
	t.Logf("Monitor URL: %s", monitor.URL)
	if len(monitor.ResponseTimes) > 0 {
		t.Logf("Monitor Recent Response Time(msec): %d", monitor.ResponseTimes[0].Value)
	}
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
