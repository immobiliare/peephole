package mold

import (
	"testing"
)

func TestMissingRetention(t *testing.T) {
	c := Config{}
	if err := c.Validate(); err == nil {
		t.Errorf("Missing retention field is supposed to return an error")
	}
}

func TestMalformedRetention(t *testing.T) {
	c := Config{Retention: "1"}
	if err := c.Validate(); err == nil {
		t.Errorf("Malformed retention field is supposed to return an error")
	}
}

func TestDefaultBind(t *testing.T) {
	c := Config{Retention: "1d"}
	if err := c.Validate(); err != nil {
		t.Errorf("Valid config fields are not supposed to return an error")
	}
	if c.Spool == "" {
		t.Errorf("Empty bind field is supposed to lead to the default value")
	}
}
