package salt

import (
	"fmt"
	"testing"
)

func TestRetryEvent(t *testing.T) {
	e, err := unmarshal([]string{"retry:"})
	if err != nil {
		t.Errorf("Unexpected failure while unmarshaling valid EventsResponse")
	}

	if e.Type != eventRetry {
		t.Errorf("Parsed event type is supposed to be retry")
	}
}

func TestDataEvent(t *testing.T) {
	var (
		tag    = "tag"
		data   = "{}"
		e, err = unmarshal([]string{
			fmt.Sprintf("tag: %s", tag),
			fmt.Sprintf(`data: {"tag": "%s", "data": %s}`, tag, data),
		})
	)
	if err != nil {
		t.Errorf("Unexpected failure while unmarshaling valid EventsResponse")
	}

	cond := e.Type == eventData
	cond = cond && e.Tag == tag
	cond = cond && e.Data == data
	if !cond {
		t.Errorf("Parsed event mismatches with input data")
	}
}

func TestUnknownEvent(t *testing.T) {
	e, err := unmarshal([]string{})
	if err != nil {
		t.Errorf("Unexpected failure while unmarshaling valid EventsResponse")
	}

	if e.Type != eventUnknown {
		t.Errorf("Parsed event type is supposed to be unknown")
	}
}
