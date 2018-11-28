package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseHostNewAPI = "http://tl-apps.t-l.ch"
	baseHost       = "http://syn.t-l.ch"
)

// Client is a http client wrapper
type Client struct {
	*http.Client
}

// NewClient creates a fershly new instance of client
func NewClient() *Client {
	return &Client{&http.Client{}}
}

func request(endpoint string) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", baseHost, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accepted-Encoding", "gzip")
	req.Header.Add("Content-Type", "application/json; charset=utf8")

	return req, nil
}

func (c *Client) execRequest(r *http.Request, v interface{}) error {
	resp, err := c.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}

// ListLines fetch all active lines on the server.
// This also updates the known lines of the caches if the reques
// succeeded.
func (c *Client) ListLines() ([]Line, error) {

	req, err := request("apps/LinesList")
	wrapping := struct {
		Lines struct {
			Line []Line `json:"line"`
		} `json:"lines"`
	}{}

	err = c.execRequest(req, &wrapping)
	return wrapping.Lines.Line, err
}

// ListStops fetch all active stops on the server.
func (c *Client) ListStops() ([]Stop, error) {

	req, err := request("apps/StopAreasList")
	wrapping := struct {
		Stops struct {
			Stop []Stop `json:"stopArea"`
		} `json:"stopAreas"`
	}{}
	err = c.execRequest(req, &wrapping)
	return wrapping.Stops.Stop, err
}

func (c *Client) ListRoutes(line Line) ([]Route, error) {
	return c.ListRoutesFromID(line.ID)
}

func (c *Client) ListRoutesFromID(ID string) ([]Route, error) {
	url := fmt.Sprintf("apps/StopAreasList?roid=%s", ID)
	req, err := request(url)

	wrapping := struct {
		Routes struct {
			Routes []Route `json:"routes"`
		} `json:"routes"`
	}{}
	err = c.execRequest(req, &wrapping)
	return wrapping.Routes.Routes, err
}
