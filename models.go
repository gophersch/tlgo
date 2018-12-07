package tlgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Message represents a message information
type Message struct {
	Content string `json:"content,omitempty"`
}

type lineRequest struct {
	Lines struct {
		Line []Line `json:"line"`
	} `json:"lines"`
}

// Line represents a line informations
type Line struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name"`
	ShortName string    `json:"line_short_name"`
	Message   []Message `json:"message"`
}

type stopRequest struct {
	Stops struct {
		Stop []Stop `json:"stopArea"`
	} `json:"stopAreas"`
}

// Stop represents stops information
type Stop struct {
	ID             string
	Name           string
	ShortName      string
	Lat            float32
	Lng            float32
	LinesShortName []string
}

func (s *Stop) UnmarshalJSON(b []byte) error {

	empty := struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"line_short_name"`
		X         string `json:"x"`
		Y         string `json:"y"`
		Lines     []struct {
			ShortName string `json:"line_short_name"`
		} `json:"line"`
	}{}

	err := json.Unmarshal(b, &empty)
	if err != nil {
		return err
	}

	s.ID = empty.ID
	s.Name = empty.Name
	s.ShortName = empty.ShortName

	s.LinesShortName = make([]string, len(empty.Lines))
	for i := range empty.Lines {
		s.LinesShortName[i] = empty.Lines[i].ShortName
	}

	lat, err := strconv.ParseFloat(empty.Y, 32)
	if err != nil {
		return fmt.Errorf("invalid latitude format: %v", err)
	}

	lng, err := strconv.ParseFloat(empty.X, 32)
	if err != nil {
		return fmt.Errorf("invalid longitude format: %v", err)
	}

	s.Lat = float32(lat)
	s.Lng = float32(lng)
	return nil
}

type routeRequest struct {
	Routes struct {
		Routes []Route `json:"routes"`
	} `json:"routes"`
}

// Route represents a route
type Route struct {
	CityDestination         string
	CityDestinationStopName string
	Direction               string
	MainRoute               bool
	Length                  float32
	Name                    string
	CityOrigin              string
	CityOriginStopName      string
	Rank                    int
	RankOdd                 bool
	ID                      string
	StopsCount              int
	Wayback                 bool
}

func (r *Route) UnmarshalJSON(b []byte) error {

	empty := struct {
		CityDestination         string `json:"destination_city_name"`
		CityDestinationStopName string `json:"destination_stop_name"`
		Direction               string `json:"direction"`
		MainRoute               string `json:"is_main"`
		Length                  string `json:"length"`
		Name                    string `json:"name"`
		CityOrigin              string `json:"origin_city_name"`
		CityOriginStopName      string `json:"origin_stop_name"`
		Rank                    string `json:"rank"`
		RankOdd                 string `json:"rank_is_odd"`
		ID                      string `json:"roid"`
		StopsCount              string `json:"stops_number"`
		Wayback                 string `json:"wayback"`
	}{}

	err := json.Unmarshal(b, &empty)
	if err != nil {
		return err
	}

	length, err := strconv.ParseFloat(empty.Length, 32)
	if err != nil {
		return nil
	}

	rank, err := strconv.ParseInt(empty.Rank, 10, 32)
	if err != nil {
		return nil
	}
	r.Rank = int(rank)
	count, err := strconv.ParseInt(empty.StopsCount, 10, 32)
	if err != nil {
		return nil
	}
	r.StopsCount = int(count)

	r.Length = float32(length)
	r.CityDestination = empty.CityDestination
	r.CityDestinationStopName = empty.CityDestinationStopName
	r.Direction = empty.Direction
	r.MainRoute = boolFromString(empty.MainRoute)
	r.Name = empty.Name
	r.CityOrigin = empty.CityOrigin
	r.CityOriginStopName = empty.CityOriginStopName

	r.RankOdd = boolFromString(empty.RankOdd)
	r.ID = empty.ID

	r.Wayback = boolFromString(empty.Wayback)

	return nil
}

func stringFromBool(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func boolFromString(s string) bool {
	if s == "1" {
		return true
	}
	return false
}

// StopRouteDetails hold info from stop inside a
// RouteDetails.
type StopRouteDetails struct {
	ID           string `json:"id"`
	StopAreaName string `json:"stopAreaName"`
}

// RouteDetails gives information about the route
type RouteDetails struct {
	LineID    string
	ShortName string
	Stops     []StopRouteDetails
	Wayback   bool
}

func (r *RouteDetails) UnmarshalJSON(b []byte) error {

	empty := struct {
		LineID    string             `json:"lineId"`
		ShortName string             `json:"lineShortName"`
		Stops     []StopRouteDetails `json:"stop"`
		Wayback   string             `json:"wayback"`
	}{}

	err := json.Unmarshal(b, &empty)
	if err != nil {
		return err
	}

	r.LineID = empty.LineID
	r.ShortName = empty.ShortName
	r.Stops = empty.Stops
	r.Wayback = boolFromString(empty.Wayback)

	return nil
}

type routeDetailsRequest struct {
	RouteDetails RouteDetails `json:"route"`
}

// JourneyLine holds line information in a journey
type JourneyLine struct {
	ID        string `json:"id"`
	ShortName string `json:"line_short_name"`
}

type JourneyStop struct {
	Name string `json:"name"`
}

// Journey holds response from next departures
type Journey struct {
	Time             time.Time
	NetworkID        string
	NetworkName      string
	RouteID          string
	Track            string
	WaitingTime      time.Duration
	DisabilityAccess bool
	Realtime         bool
	Stops            []JourneyStop
	Wayback          bool
	Message          []Message
	Lines            []JourneyLine
}

func (j *Journey) UnmarshalJSON(b []byte) error {

	empty := struct {
		Time             string        `json:"date_time"`
		DisabilityAccess string        `json:"handicapped_access"`
		Realtime         string        `json:"realTime"`
		RouteID          string        `json:"route_id"`
		Track            string        `json:"track"`
		Wayback          string        `json:"wayback"`
		Stops            []JourneyStop `json:"stop"`
		Message          []Message     `json:"message"`
		Lines            []JourneyLine `json:"line"`
		NetworkID        string        `json:"networkId"`
		NetworkName      string        `json:"networkName"`
		WaitingTime      string        `json:"waiting_time"`
	}{}

	err := json.Unmarshal(b, &empty)
	if err != nil {
		return err
	}

	timeValue, err := time.Parse("2006-01-02 15:04:05", empty.Time)
	if err != nil {
		return err
	}
	j.Time = timeValue

	j.Stops = empty.Stops

	j.DisabilityAccess = boolFromString(empty.DisabilityAccess)
	j.Realtime = boolFromString(empty.Realtime)
	j.RouteID = empty.RouteID
	j.Track = empty.Track
	j.NetworkID = empty.NetworkID
	j.NetworkName = empty.NetworkName

	j.Wayback = boolFromString(empty.Wayback)
	j.Message = empty.Message
	j.Lines = empty.Lines

	// Waiting time
	if len(empty.WaitingTime) != 8 {
		return errors.New("Invalid waiting time format")
	}

	hours, hoursErr := strconv.ParseInt(empty.WaitingTime[0:2], 10, 8)
	minutes, minutesErr := strconv.ParseInt(empty.WaitingTime[3:5], 10, 8)
	seconds, secondsErr := strconv.ParseInt(empty.WaitingTime[6:8], 10, 8)
	if hoursErr != nil || minutesErr != nil || secondsErr != nil {
		return errors.New("Can not parse the waiting time")
	}

	j.WaitingTime = time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second

	return nil
}

type journeyRequest struct {
	Journeys struct {
		Journey []Journey `json:"journey"`
	} `json:"journeys"`
}
