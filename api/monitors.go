package api

import (
	"errors"
	"strconv"
)

type MonitorType int

const (
	Http MonitorType = 1 + iota
	Keyword
	Ping
	Port
)

type NewMonitorRequest struct {
	FriendlyName string
	Url          string
	MonitorType  MonitorType
}

type EditMonitorRequest struct {
	Id           int
	FriendlyName string
	Url          string
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
	ID               int                 `xml:"id,int,attr"`
	FriendlyName     string              `xml:"friendly_name,string,attr"`
	URL              string              `xml:"url,string,attr"`
	ResponseTimeList XMLResponseTimeList `xml:"response_times"`
}

type XMLResponseTimeList struct {
	ResponseTimes []ResponseTime `xml:"response_time"`
}

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

func (r *request) setDeleteMonitorRequest(req DeleteMonitorRequest) error {
	if req.Id == 0 {
		return errors.New("Id: required value")
	}
	r.body.Set("id", strconv.Itoa(req.Id))
	return nil
}

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

func (r *request) setGetMonitorsRequest(req GetMonitorsRequest) error {
	if req.MonitorId == 0 {
		return errors.New("MonitorId: required value")
	}
	r.body.Set("monitors", strconv.Itoa(req.MonitorId))
	r.body.Set("response_times", "1")
	r.body.Set("response_times_average", "300")
	return nil
}
