package tlgo

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

func (c *Client) execRequest(enpoint string, v interface{}) error {
	r, err := request(enpoint)
	if err != nil {
		return err
	}

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

	wrapping := lineRequest{}
	err := c.execRequest("apps/LinesList", &wrapping)
	return wrapping.Lines.Line, err
}

// ListStops fetch all active stops on the server.
func (c *Client) ListStops() ([]Stop, error) {

	wrapping := stopRequest{}
	err := c.execRequest("apps/StopAreasList", &wrapping)
	return wrapping.Stops.Stop, err
}

// ListRoutes list the routes contained in a `Line`
func (c *Client) ListRoutes(line *Line) ([]Route, error) {
	return c.ListRoutesFromID(line.ID)
}

// ListRoutesFromID list the routes contained in a `Line`
// from the line ID.
func (c *Client) ListRoutesFromID(ID string) ([]Route, error) {
	url := fmt.Sprintf("apps/RoutesLists?lineid=%s", ID)
	wrapping := routeRequest{}
	err := c.execRequest(url, &wrapping)
	return wrapping.Routes.Routes, err
}
