package api

import (
	"errors"
	"strconv"
)

// AlertContactType defines the type for an alert contact
type AlertContactType int

const (
	AlertContactSMS AlertContactType = 1 + iota
	AlertContactEmail
	AlertContactTwitterDM
	AlertContactBoxcar
	AlertContactWebHook
	AlertContactPushbullet
	AlertContactZapier
	AlertContactPushover
	AlertContactHipChat
	AlertContactSlack
)

// AlertContacts is used to access the UptimeRobot alert contacts
type AlertContacts struct {
	c *Client
}

// AlertContacts is used to return a handle to the alert contacts apis
func (c *Client) AlertContacts() *AlertContacts {
	return &AlertContacts{c}
}

// AlertContactResponse contains the response from the server for containing an alert contact
type AlertContactResponse struct {
	ID int `xml:"id,int,attr"`
}

// NewAlertContactRequest contains parameters for a new alert contact request
type NewAlertContactRequest struct {
	AlertContactType         AlertContactType
	AlertContactValue        string
	AlertContactFriendlyName string
}

// DeleteAlertContactRequest contains parameters for a delete alert contact request
type DeleteAlertContactRequest struct {
	ID int
}

// GetAlertContactsResponse contains a reponse for get alert contacts request
type GetAlertContactsResponse struct {
	AlertContacts []alertContact `xml:"alertcontact"`
}

type alertContact struct {
	ID           int              `xml:"id,int,attr"`
	Value        string           `xml:"value,string,attr"`
	FriendlyName string           `xml:"friendlyname,string,attr"`
	Type         AlertContactType `xml:"type,int,attr"`
	Status       int              `xml:"status,int,attr"`
}

// New creates a new alert contact
func (ac *AlertContacts) New(req NewAlertContactRequest) (*AlertContactResponse, error) {
	r := ac.c.newRequest("GET", "/newAlertContact")
	err := r.setNewAlertContactRequest(req)
	if err != nil {
		return nil, err
	}

	_, resp, err := requireOK(ac.c.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out *AlertContactResponse
	if err := decodeBody(resp, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *request) setNewAlertContactRequest(req NewAlertContactRequest) error {
	if req.AlertContactType == 0 {
		return errors.New("alertContactType: required value")
	}
	if req.AlertContactValue == "" {
		return errors.New("AlertContactValue: required value")
	}
	r.params.Set("alertContactType", strconv.Itoa(int(req.AlertContactType)))
	r.params.Set("alertContactValue", req.AlertContactValue)
	if req.AlertContactFriendlyName != "" {
		r.params.Set("alertContactFriendlyName", req.AlertContactFriendlyName)
	}

	return nil
}

// Delete deletes an alert contact
func (ac *AlertContacts) Delete(req DeleteAlertContactRequest) (*AlertContactResponse, error) {
	r := ac.c.newRequest("GET", "/deleteAlertContact")
	err := r.setDeleteAlertContactRequest(req)
	if err != nil {
		return nil, err
	}

	_, resp, err := requireOK(ac.c.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out *AlertContactResponse
	if err := decodeBody(resp, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *request) setDeleteAlertContactRequest(req DeleteAlertContactRequest) error {
	if req.ID == 0 {
		return errors.New("Id: required value")
	}
	r.params.Set("alertContactID", strconv.Itoa(req.ID))
	return nil
}

// Get retrieves a list of alert contacts
func (ac *AlertContacts) Get() (*GetAlertContactsResponse, error) {
	r := ac.c.newRequest("GET", "/getAlertContacts")

	_, resp, err := requireOK(ac.c.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out *GetAlertContactsResponse
	if err := decodeBody(resp, &out); err != nil {
		return nil, err
	}

	return out, nil
}
