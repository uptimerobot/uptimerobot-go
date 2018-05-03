package api

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

// Account is UptimeRobot account representation
type Account struct {
	MonitorLimit    int `xml:"monitor_limit,int,attr"`
	MonitorInterval int `xml:"monitor_interval,int,attr"`
	UpMonitors      int `xml:"up_monitors,int,attr"`
	DownMonitors    int `xml:"down_monitors,int,attr"`
	PausedMonitors  int `xml:"paused_monitors,int,attr"`
}

// AccountDetails is used to access the UptimeRobot account details
type AccountDetails struct {
	c *Client
}

// AccountDetails is used to return a handle to the AccountDetails apis
func (c *Client) AccountDetails() *AccountDetails {
	return &AccountDetails{c}
}

// Get calls UptimeRobot API to retrieve Account Details information
func (ad *AccountDetails) Get() (*Account, error) {
	r := ad.c.newRequest("POST", "/v2/getAccountDetails")
	duration, resp, err := ad.c.doRequest(r)
	if err != nil {
		return nil, err
	}
	_, resp, err = requireOK(duration, resp, err)
	if err != nil {
		log.Print("is not ok?")
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var out *Account
	if err := xml.Unmarshal(body, &out); err != nil {
		return nil, err
	}
	return out, nil
}
