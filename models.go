package apiclient

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Message represents a message information
type Message struct {
	Content string `json:"content,omitempty"`
}

// Line represents a line informations
type Line struct {
	ID        string  `json:"id,omitempty"`
	Name      string  `json:"name"`
	ShortName string  `json:"line_short_name"`
	Message   Message `json:"message"`
}

// Stop represents stops information
type Stop struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	ShortName string  `json:"line_short_name"`
	Lat       float32 `json:"y"`
	Lng       float32 `json:"x"`
	Lines     []Line  `json:"line"`
}

func (s *Stop) UnmarshalJSON(b []byte) error {

	empty := struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"line_short_name"`
		X         string `json:"x"`
		Y         string `json:"y"`
		Lines     []Line `json:"line"`
	}{}

	err := json.Unmarshal(b, &empty)
	if err != nil {
		return err
	}

	s.ID = empty.ID
	s.Name = empty.Name
	s.ShortName = empty.ShortName
	s.Lines = empty.Lines

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

type Route struct {
	CityDestination         string  `json:"destination_city_name"`
	CityDestinationStopName string  `json:"destination_stop_name"`
	Direction               string  `json:"direction"`
	MainRoute               bool    `json:"is_main"`
	Length                  float32 `json:"length"`
	Name                    string  `json:"name"`
	CityOrigin              string  `json:"origin_city_name"`
	CityOriginStopName      string  `json:"origin_stop_name"`
	Rank                    int     `json:"rank"`
	RankOdd                 bool    `json:"rank_is_odd"`
	ID                      string  `json:"roid"`
	StopsCount              int     `json:"stops_number"`
	Wayback                 bool    `json:"wayback"`
}
