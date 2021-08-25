package spy

import (
	"testing"
)

func TestConfig(t *testing.T) {
	c := Config{}
	if err := c.Validate(); err != nil {
		t.Errorf("Validation is not supposed to return an error")
	}
}
