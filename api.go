package tlgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	baseHost        = "http://syn.t-l.ch"
	timeQueryLayout = "2006-01-02 15:04"
)

// Client is a http client wrapper
type Client struct {
	http *http.Client
}

// NewClient creates a fershly new instance of client
func NewClient() *Client {
	return &Client{
		http: &http.Client{},
	}
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

	resp, err := c.http.Do(r)
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
// from the line ID.
func (c *Client) ListRoutes(ID string) ([]Route, error) {
	url := fmt.Sprintf("apps/RoutesList?lineid=%s", ID)
	wrapping := routeRequest{}
	err := c.execRequest(url, &wrapping)
	return wrapping.Routes.Routes, err
}

// GetRouteDetails returns the list of a route.
func (c *Client) GetRouteDetails(ID string) (RouteDetails, error) {
	url := fmt.Sprintf("apps/RouteDetail?roid=%s", ID)
	wrapping := routeDetailsRequest{}
	err := c.execRequest(url, &wrapping)
	return wrapping.RouteDetails, err
}

// ListStopDepartures retrieve the nest departures informations for a line
func (c *Client) ListStopDepartures(stopID string, lineID string, date time.Time, wayback bool) ([]Journey, error) {

	URL := departurePath(stopID, lineID, date, wayback)
	wrapping := journeyRequest{}
	err := c.execRequest(URL, &wrapping)
	return wrapping.Journeys.Journey, err
}

func departurePath(stopID, lineID string, date time.Time, wayback bool) string {
	return fmt.Sprintf("apps/LineStopDeparturesList?date=%s&lineid=%s&roid=%s&wayback=%s", url.QueryEscape(date.Format(timeQueryLayout)), lineID, stopID, stringFromBool(wayback))
}
