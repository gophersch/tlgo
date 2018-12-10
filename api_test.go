package tlgo

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func validJSON(v interface{}, filename string, t *testing.T) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Can not open sample file: %s", err)
		t.FailNow()
	}
	err = json.NewDecoder(strings.NewReader(string(file))).Decode(v)
	if err != nil {
		t.Errorf("Can not parse %s correctly: %s", reflect.TypeOf(v), err)
	}

	return err
}
func TestJSON(t *testing.T) {
	lineRequest := &lineRequest{}
	validJSON(lineRequest, "samples/line_request.json", t)
	assert.Equal(t, len(lineRequest.Lines.Line), 2)

	first := lineRequest.Lines.Line[0]

	assert.Equal(t, first.ID, "11822125115506799")
	assert.Equal(t, first.Name, "Lausanne-Flon - Bercher")
	assert.Equal(t, first.ShortName, "LEB")
	assert.Equal(t, len(first.Message), 1)

	validJSON(&stopRequest{}, "samples/line_request.json", t)

	journey := journeyRequest{}
	validJSON(&journey, "samples/next_departure.json", t)

	assert.Equal(t, len(journey.Journeys.Journey), 12)

	j := journey.Journeys.Journey[0]
	expectedDuration := time.Duration(-11)*time.Minute - time.Duration(7)*time.Second
	assert.Equal(t, j.WaitingTime, expectedDuration)

	j2 := journey.Journeys.Journey[1]
	expectedDuration = time.Duration(12)*time.Minute + time.Duration(27)*time.Second
	assert.Equal(t, j2.WaitingTime, expectedDuration)

}

func TestURLMarshalling(t *testing.T) {

	date := time.Date(2018, time.December, 7, 14, 22, 0, 0, time.UTC)
	expectedPath := "apps/LineStopDeparturesList?date=2018-12-07+14%3A22&lineid=11821953316814862&roid=1970329131942119&wayback=1"

	URL := departurePath("1970329131942119", "11821953316814862", date, true)
	if URL != expectedPath {
		t.Errorf("URL is not the expected one: %s != %s", expectedPath, URL)
	}

}
