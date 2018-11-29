package tlgo

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

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
}
