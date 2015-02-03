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
	friendlyName string
	url          string
	monitorType  MonitorType
}

type EditMonitorRequest struct {
	id           int
	friendlyName string
	url          string
	monitorType  MonitorType
}

type DeleteMonitorRequest struct {
	id int
}

type MonitorResponse struct {
	ID int `xml:"id,int,attr"`
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

func (r *request) setNewMonitorRequest(req NewMonitorRequest) error {
	if req.friendlyName == "" {
		return errors.New("friendlyName: required value")
	}
	if req.url == "" {
		return errors.New("url: required value")
	}
	r.params.Set("monitorFriendlyName", req.friendlyName)
	r.params.Set("monitorURL", req.url)
	r.params.Set("monitorType", strconv.Itoa(int(req.monitorType)))
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
	if req.id == 0 {
		return errors.New("id: required value")
	}
	r.params.Set("monitorID", strconv.Itoa(req.id))
	if req.friendlyName != "" {
		r.params.Set("monitorFriendlyName", req.friendlyName)
	}
	if req.url != "" {
		r.params.Set("monitorURL", req.url)
	}
	if int(req.monitorType) != 0 {
		r.params.Set("monitorType", strconv.Itoa(int(req.monitorType)))
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
	if req.id == 0 {
		return errors.New("id: required value")
	}
	r.params.Set("monitorID", strconv.Itoa(req.id))
	return nil
}
