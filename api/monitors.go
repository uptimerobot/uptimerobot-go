package api

import (
	"errors"
	"strconv"
	"strings"
)

type MonitorType int

const (
	Http MonitorType = 1 + iota
	Keyword
	Ping
	Port
)

type NewMonitorRequest struct {
	FriendlyName    string
	Url             string
	MonitorType     MonitorType
	AlertContactIDs []int
}

type EditMonitorRequest struct {
	Id              int
	FriendlyName    string
	Url             string
	MonitorType     MonitorType
	AlertContactIDs []int
}

type DeleteMonitorRequest struct {
	Id int
}

type MonitorResponse struct {
	ID int `xml:"id,int,attr"`
}

type GetMonitorsRequest struct {
	MonitorId int
}

type XMLMonitors struct {
	Monitors []XMLMonitor `xml:"monitor"`
}

type XMLMonitor struct {
	ID            int               `xml:"id,int,attr"`
	FriendlyName  string            `xml:"friendlyname,string,attr"`
	URL           string            `xml:"url,string,attr"`
	ResponseTimes []XMLResponseTime `xml:"responsetime"`
	AlertContacts []alertContact    `xml:"alertcontact"`
}

type XMLResponseTime struct {
	Value int `xml:"value,int,attr"`
}

// Monitors is used to access the UptimeRobot monitors
type Monitors struct {
	c *Client
}

// Monitors is used to return a handle to the monitors apis
func (c *Client) Monitors() *Monitors {
	return &Monitors{c}
}

func (ad *Monitors) New(req NewMonitorRequest) (*MonitorResponse, error) {
	r := ad.c.newRequest("GET", "/newMonitor")
	err := r.setNewMonitorRequest(req)
	if err != nil {
		return nil, err
	}

	_, resp, err := requireOK(ad.c.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out *MonitorResponse
	if err := decodeBody(resp, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func formatAlertContactIDs(alertContactIDs []int) string {
	alertContactIDsAsStrings := make([]string, len(alertContactIDs))
	for idx, id := range alertContactIDs {
		alertContactIDsAsStrings[idx] = strconv.Itoa(id) + "_0_0"
	}
	return strings.Join(alertContactIDsAsStrings, "-")
}

func (r *request) setNewMonitorRequest(req NewMonitorRequest) error {
	if req.FriendlyName == "" {
		return errors.New("FriendlyName: required value")
	}
	if req.Url == "" {
		return errors.New("Url: required value")
	}
	r.params.Set("monitorFriendlyName", req.FriendlyName)
	r.params.Set("monitorURL", req.Url)
	r.params.Set("monitorType", strconv.Itoa(int(req.MonitorType)))
	if len(req.AlertContactIDs) > 0 {
		r.params.Set("monitorAlertContacts", formatAlertContactIDs(req.AlertContactIDs))
	}
	return nil
}

func (ad *Monitors) Edit(req EditMonitorRequest) (*MonitorResponse, error) {
	r := ad.c.newRequest("GET", "/editMonitor")
	err := r.setEditMonitorRequest(req)
	if err != nil {
		return nil, err
	}

	_, resp, err := requireOK(ad.c.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out *MonitorResponse
	if err := decodeBody(resp, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *request) setEditMonitorRequest(req EditMonitorRequest) error {
	if req.Id == 0 {
		return errors.New("Id: required value")
	}
	r.params.Set("monitorID", strconv.Itoa(req.Id))
	if req.FriendlyName != "" {
		r.params.Set("monitorFriendlyName", req.FriendlyName)
	}
	if req.Url != "" {
		r.params.Set("monitorURL", req.Url)
	}
	if int(req.MonitorType) != 0 {
		r.params.Set("monitorType", strconv.Itoa(int(req.MonitorType)))
	}
	if len(req.AlertContactIDs) > 0 {
		r.params.Set("monitorAlertContacts", formatAlertContactIDs(req.AlertContactIDs))
	}
	return nil
}

func (ad *Monitors) Delete(req DeleteMonitorRequest) (*MonitorResponse, error) {
	r := ad.c.newRequest("GET", "/deleteMonitor")
	err := r.setDeleteMonitorRequest(req)
	if err != nil {
		return nil, err
	}

	_, resp, err := requireOK(ad.c.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out *MonitorResponse
	if err := decodeBody(resp, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *request) setDeleteMonitorRequest(req DeleteMonitorRequest) error {
	if req.Id == 0 {
		return errors.New("Id: required value")
	}
	r.params.Set("monitorID", strconv.Itoa(req.Id))
	return nil
}

func (ad *Monitors) Get(req GetMonitorsRequest) (*XMLMonitors, error) {
	r := ad.c.newRequest("GET", "/getMonitors")
	err := r.setGetMonitorsRequest(req)
	if err != nil {
		return nil, err
	}

	_, resp, err := requireOK(ad.c.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out *XMLMonitors
	if err := decodeBody(resp, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (r *request) setGetMonitorsRequest(req GetMonitorsRequest) error {
	if req.MonitorId == 0 {
		return errors.New("monitors: required value")
	}
	r.params.Set("monitors", strconv.Itoa(req.MonitorId))
	r.params.Set("responseTimes", "1")
	r.params.Set("responseTimesAverage", "300")
	return nil
}
