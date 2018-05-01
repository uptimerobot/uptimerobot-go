package api

import (
	"errors"
	"strconv"
)

// MonitorType typo of monitor
type MonitorType int

const (
	// HTTP Monitor type
	HTTP MonitorType = 1 + iota
	// Keyword Monitor type
	Keyword
	// Ping Monitor type
	Ping
	// Port Monitor type
	Port
)

// NewMonitorRequest struct for New Monitor
type NewMonitorRequest struct {
	FriendlyName string
	URL          string
	MonitorType  MonitorType
}

// EditMonitorRequest struct for Edit Monitor
type EditMonitorRequest struct {
	ID           int
	FriendlyName string
	URL          string
	MonitorType  MonitorType
}

// DeleteMonitorRequest struct for Delete Monitor
type DeleteMonitorRequest struct {
	ID int
}

// MonitorResponse map UptimeRobot response
type MonitorResponse struct {
	ID int `xml:"id,int,attr"`
}

// GetMonitorsRequest Request for Monitors
type GetMonitorsRequest struct {
	MonitorID int
}

// XMLMonitors XML response with list monitors
type XMLMonitors struct {
	Monitors []XMLMonitor `xml:"monitor"`
}

// XMLMonitor XML representation of Monitor
type XMLMonitor struct {
	ID            int               `xml:"id,int,attr"`
	FriendlyName  string            `xml:"friendlyname,string,attr"`
	URL           string            `xml:"url,string,attr"`
	Status        string            `xml:"status,string,attr"`
	Type          string            `xml:"type,string,attr"`
	SubType       string            `xml:"sub_type,string,attr"`
	ResponseTimes []XMLResponseTime `xml:"responsetime"`
}

// XMLResponseTime XML representation of Monitor Response Time
type XMLResponseTime struct {
	Value int `xml:"value,int,attr"`
}

// Monitors is used to return a handle to the monitors apis
func (c *Client) Monitors() *Monitors {
	return &Monitors{c}
}

// Monitors is used to access the UptimeRobot monitors
type Monitors struct {
	c *Client
}

// New Request for creating a new Monitor
// See NewMonitorRequest
func (ad *Monitors) New(req NewMonitorRequest) (*MonitorResponse, error) {
	r := ad.c.newRequest("POST", "/v2/newMonitor")
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

func (r *request) setNewMonitorRequest(req NewMonitorRequest) error {
	if req.FriendlyName == "" {
		return errors.New("FriendlyName: required value")
	}
	if req.URL == "" {
		return errors.New("URL: required value")
	}
	r.params.Set("friendly_name", req.FriendlyName)
	r.params.Set("url", req.URL)
	r.params.Set("type", strconv.Itoa(int(req.MonitorType)))
	return nil
}

// Edit and existing Monitor
// EditMonitorRequest
func (ad *Monitors) Edit(req EditMonitorRequest) (*MonitorResponse, error) {
	r := ad.c.newRequest("POST", "/v2/editMonitor")
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
	if req.ID == 0 {
		return errors.New("ID: required value")
	}
	r.params.Set("id", strconv.Itoa(req.ID))
	if req.FriendlyName != "" {
		r.params.Set("friendly_ame", req.FriendlyName)
	}
	if req.URL != "" {
		r.params.Set("url", req.URL)
	}
	if int(req.MonitorType) != 0 {
		r.params.Set("type", strconv.Itoa(int(req.MonitorType)))
	}
	return nil
}

// Delete and existing Monitor
// See DeleteMonitorRequest
func (ad *Monitors) Delete(req DeleteMonitorRequest) (*MonitorResponse, error) {
	r := ad.c.newRequest("POST", "/v2/deleteMonitor")
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
	if req.ID == 0 {
		return errors.New("ID: required value")
	}
	r.params.Set("id", strconv.Itoa(req.ID))
	return nil
}

// Get and existing Monitor.
// See: GetMonitorsRequest
func (ad *Monitors) Get(req GetMonitorsRequest) (*XMLMonitors, error) {
	r := ad.c.newRequest("POST", "/v2/getMonitors")
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
	if req.MonitorID == 0 {
		return errors.New("monitors: required value")
	}
	r.params.Set("monitors", strconv.Itoa(req.MonitorID))
	r.params.Set("responseTimes", "1")
	r.params.Set("responseTimesAverage", "300")
	return nil
}
