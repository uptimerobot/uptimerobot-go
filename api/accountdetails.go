package api

// Account is used to construct and return Account Details
// For more information on the provided data, see https://uptimerobot.com/api/#getAccountDetailsWrap
type Account struct {
	Email           string `xml:"email,string,attr"`
	MonitorLimit    int    `xml:"monitor_limit,int,attr"`
	MonitorInterval int    `xml:"monitor_interval,int,attr"`
	UpMonitors      int    `xml:"up_monitors,int,attr"`
	DownMonitors    int    `xml:"down_monitors,int,attr"`
	PausedMonitors  int    `xml:"paused_monitors,int,attr"`
}

// AccountDetails is used to access the UptimeRobot account details
type AccountDetails struct {
	c *Client
}

// AccountDetails is used to return a handle to the AccountDetails apis
func (c *Client) AccountDetails() *AccountDetails {
	return &AccountDetails{c}
}

// Returns Account Details for the current client
func (ad *AccountDetails) Get() (*Account, error) {
	r := ad.c.newRequest("POST", "/getAccountDetails")
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
