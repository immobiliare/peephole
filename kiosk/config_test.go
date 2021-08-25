package kiosk

import (
	"testing"
)

func TestMalformedBind(t *testing.T) {
	c := Config{Bind: "8080"}
	if err := c.Validate(); err == nil {
		t.Errorf("Malformed bind field is supposed to return an error")
	}
}
