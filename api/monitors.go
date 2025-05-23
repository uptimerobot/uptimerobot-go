package api

import (
	"errors"
	"strconv"
)

// Allows mapping human readable words to the integers uptimerobot expects for monitoring types
// i.e. http, keyword, ping, port
type MonitorType int

const (
	Http MonitorType = 1 + iota
	Keyword
	Ping
	Port
)

// NewMonitorRequest provides the data to create a new Monitor
type NewMonitorRequest struct {
	// Friendly name for the monitor
	FriendlyName string
	// Url to monitor
	Url string
	// Type of monitoring to use
	MonitorType MonitorType
}

// EditMonitorRequest provides the data to edit a specified Monitor
type EditMonitorRequest struct {
	// Id of the monitor to edit
	Id int
	// New friendly name
	FriendlyName string
	// New url
	Url string
}

// DeleteMonitorRequest provides the data to delete a specified Monitor
type DeleteMonitorRequest struct {
	Id int
}

// MonitorResponse contains an ID for a monitor
type MonitorResponse struct {
	ID int `xml:"id,int,attr"`
}

// GetMonitorsRequest provides the data to request Monitor information
type GetMonitorsRequest struct {
	MonitorId int
}

// XMLMonitors contains a slice of XMLMonitor structs
type XMLMonitors struct {
	Monitors []XMLMonitor `xml:"monitor"`
}

// XML Monitor is used to construct and return details for one monitor
type XMLMonitor struct {
	ID               int                 `xml:"id,int,attr"`
	FriendlyName     string              `xml:"friendly_name,string,attr"`
	URL              string              `xml:"url,string,attr"`
	ResponseTimeList XMLResponseTimeList `xml:"response_times"`
}

// XMLResponseTimeList contains a slice of ResponseTime structs
type XMLResponseTimeList struct {
	ResponseTimes []ResponseTime `xml:"response_time"`
}

// ResponseTime is used to parse the response time data for a monitor
type ResponseTime struct {
	Date  int `xml:"datetime,int,attr"`
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

// Returns a MonitorResponse, containing an ID for the monitor edited,
// or an error.
func (ad *Monitors) New(req NewMonitorRequest) (*MonitorResponse, error) {
	r := ad.c.newRequest("POST", "/newMonitor")
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

// Helper func for New to construct http request body
// using the provided NewMonitorRequest struct
func (r *request) setNewMonitorRequest(req NewMonitorRequest) error {
	if req.FriendlyName == "" {
		return errors.New("FriendlyName: required value")
	}
	if req.Url == "" {
		return errors.New("Url: required value")
	}
	r.body.Set("friendly_name", req.FriendlyName)
	r.body.Set("url", req.Url)
	r.body.Set("type", strconv.Itoa(int(req.MonitorType)))
	return nil
}

// Returns a MonitorResponse, containing an ID for the monitor edited,
// or an error.
// Per https://uptimerobot.com/api/#editMonitorWrap, monitor type cannot be edited.
// To change type, monitors should be deleted and recreated.
func (ad *Monitors) Edit(req EditMonitorRequest) (*MonitorResponse, error) {
	r := ad.c.newRequest("POST", "/editMonitor")
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

// Helper func for Edit to construct http request body
// using the provided EditMonitorRequest struct
func (r *request) setEditMonitorRequest(req EditMonitorRequest) error {
	if req.Id == 0 {
		return errors.New("Id: required value")
	}
	r.body.Set("id", strconv.Itoa(req.Id))
	if req.FriendlyName != "" {
		r.body.Set("friendly_name", req.FriendlyName)
	}
	if req.Url != "" {
		r.body.Set("url", req.Url)
	}
	return nil
}

// Returns a MonitorResponse, containing an ID for the monitor deleted,
// or an error
func (ad *Monitors) Delete(req DeleteMonitorRequest) (*MonitorResponse, error) {
	r := ad.c.newRequest("POST", "/deleteMonitor")
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

// Helper func for Delete to construct http request body
// using the provided DeleteMonitorRequest struct
func (r *request) setDeleteMonitorRequest(req DeleteMonitorRequest) error {
	if req.Id == 0 {
		return errors.New("Id: required value")
	}
	r.body.Set("id", strconv.Itoa(req.Id))
	return nil
}

// Returns an XMLMonitors struct,
// containing a list of monitors with their associated data,
// or an error
func (ad *Monitors) Get(req GetMonitorsRequest) (*XMLMonitors, error) {
	r := ad.c.newRequest("POST", "/getMonitors")
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

// Helper func for Get to construct http request body
// using the provided GetMonitorsRequest struct
func (r *request) setGetMonitorsRequest(req GetMonitorsRequest) error {
	if req.MonitorId == 0 {
		return errors.New("MonitorId: required value")
	}
	r.body.Set("monitors", strconv.Itoa(req.MonitorId))
	r.body.Set("response_times", "1")
	r.body.Set("response_times_average", "300")
	return nil
}
