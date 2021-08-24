package kiosk

import (
	"fmt"
	"strings"
)

// Config represents the config struct for Kiosk module
type Config struct {
	Bind      string            `yaml:"bind"`
	BasicAuth map[string]string `yaml:"basic_auth"`
}

func (c *Config) Validate() error {
	if !strings.ContainsRune(c.Bind, ':') {
		return fmt.Errorf("malformed bind field")
	}

	return nil
}
