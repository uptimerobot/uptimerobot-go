package api

import (
	"os"
	"strconv"
	"strings"
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

	envMonitorIdList := os.Getenv("UPTIMEROBOT_MONITOR_ID_LIST") // Should be CSV with at least two IDs
	if envMonitorIdList == "" {
		t.Skip("TestGetMonitors requires UPTIMEROBOT_MONITOR_ID_LIST env variable")
	}
	monitorIds := []int{}
	for _, idStr := range strings.Split(envMonitorIdList, ",") {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			t.Fatal(err)
		}
		monitorIds = append(monitorIds, id)
	}

	c := makeClient(t)

	monitors := c.Monitors()

	getTests := []struct {
		name    string
		request GetMonitorsRequest
	}{
		{name: "Single ID", request: GetMonitorsRequest{MonitorIds: []int{monitorId}}},
		{name: "All monitors", request: GetMonitorsRequest{MonitorIds: []int{}}},
		{name: "Multiple monitors", request: GetMonitorsRequest{MonitorIds: monitorIds}},
	}
	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {

			response, err := monitors.Get(tt.request)
			if err != nil {
				t.Fatal(err)
			}
			if response == nil {
				t.Fatalf("No monitor response: %v", response)
			}
			for _, monitor := range response.Monitors {
				t.Logf("Monitor Friendly Name: %s\n", monitor.FriendlyName)
			}
		})
	}
}
