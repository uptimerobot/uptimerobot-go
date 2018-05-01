package v2

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const endpoint = "api.uptimerobot.com"

// Config is used to configure the creation of a client
type Config struct {
	// Address is the address of the UptimeRobot service
	Address string

	// HTTPClient is the client to use. Default will be
	// used if not provided.
	HTTPClient *http.Client

	// WaitTime limits how long a Watch will block. If not provided,
	// the agent default values will be used.
	WaitTime time.Duration

	// API key is used to provide a per-request authentication
	APIKey string
}

// Client provides a client to the UptimeRobot API
type Client struct {
	config Config
}

// request is used to help build up a request
type request struct {
	config *Config
	method string
	url    *url.URL
	params url.Values
	body   io.Reader
	obj    interface{}
}

// NewClient returns a new client
func NewClient(apikey string) (*Client, error) {
	config := &Config{
		Address:    endpoint,
		HTTPClient: http.DefaultClient,
		APIKey:     apikey,
	}

	client := &Client{
		config: *config,
	}
	return client, nil
}

// toHTTP converts the request to an HTTP request
func (r *request) toHTTP() (*http.Request, error) {
	// Encode the query parameters
	r.body = strings.NewReader(r.params.Encode())

	// Create the HTTP request
	req, err := http.NewRequest(r.method, r.url.RequestURI(), r.body)
	if err != nil {
		return nil, err
	}

	req.URL.Host = r.url.Host
	req.URL.Scheme = r.url.Scheme
	req.Host = r.url.Host
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (c *Client) newRequest(method, path string) *request {
	r := &request{
		config: &c.config,
		method: method,
		url: &url.URL{
			Scheme: "https",
			Host:   c.config.Address,
			Path:   path,
		},
		params: make(map[string][]string),
	}
	if c.config.APIKey != "" {
		r.params.Set("api_key", r.config.APIKey)
	}
	// Hardcode to xml output. Can't get nested json to decode correctly into nested structs.
	r.params.Set("format", "xml")
	return r
}

// doRequest runs a request with our client
func (c *Client) doRequest(r *request) (time.Duration, *http.Response, error) {
	req, err := r.toHTTP()
	if err != nil {
		return 0, nil, err
	}
	start := time.Now()
	resp, err := c.config.HTTPClient.Do(req)
	diff := time.Now().Sub(start)
	return diff, resp, err
}

// decodeBody is used to XML decode a body
func decodeBody(resp *http.Response, out interface{}) error {
	dec := xml.NewDecoder(resp.Body)
	return dec.Decode(out)
}

// requireOK is used to wrap doRequest and check for a 200
func requireOK(d time.Duration, resp *http.Response, e error) (time.Duration, *http.Response, error) {
	if e != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return d, nil, e
	}
	if resp.StatusCode != 200 {
		var buf bytes.Buffer
		io.Copy(&buf, resp.Body)
		resp.Body.Close()
		return d, nil, fmt.Errorf("Unexpected response code: %d (%s)", resp.StatusCode, buf.Bytes())
	}
	return d, resp, nil
}
