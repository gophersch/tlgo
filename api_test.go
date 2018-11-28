package apiclient

import (
	"testing"
)

func TestAllLines(t *testing.T) {

	_, err := NewClient()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

}
