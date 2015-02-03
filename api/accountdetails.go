package api

type Account struct {
	MonitorLimit   int `xml:"monitorlimit,int,attr"`
	UpMonitors     int `xml:"upmonitors,int,attr"`
	DownMonitors   int `xml:"downmonitors,int,attr"`
	PausedMonitors int `xml:"pausedmonitors,int,attr"`
}

// AccountDetails is used to access the UptimeRobot account details
type AccountDetails struct {
	c *Client
}

// AccountDetails is used to return a handle to the AccountDetails apis
func (c *Client) AccountDetails() *AccountDetails {
	return &AccountDetails{c}
}

func (ad *AccountDetails) Get() (*Account, error) {
	r := ad.c.newRequest("GET", "/getAccountDetails")
	_, resp, err := requireOK(ad.c.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out *Account
	if err := decodeBody(resp, &out); err != nil {
		return nil, err
	}
	return out, nil
}
