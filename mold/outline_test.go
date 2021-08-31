package mold

import "testing"

func TestOutline(t *testing.T) {
	e := Event{Raw: "not empty"}
	if len(e.Outline().Raw) > 0 {
		t.Errorf("Malformed event outline")
	}
}
