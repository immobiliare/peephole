package mold

import "testing"

func TestOutline(t *testing.T) {
	e := Event{RawData: "not empty"}
	if len(e.Outline().RawData) > 0 {
		t.Errorf("Malformed event outline")
	}
}
