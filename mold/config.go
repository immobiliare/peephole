package mold

import (
	"fmt"
	"strconv"
	"strings"

	_util "github.com/streambinder/peephole/util"
)

// Config represents the config struct for Mold module
type Config struct {
	Spool     string `yaml:"spool"`
	Retention string `yaml:"retention"`
}

func (c *Config) Validate() error {
	if c.Retention == "" {
		return fmt.Errorf("retention field is mandatory")
	}

	u, err := _util.Unit(c.Retention)
	if err != nil {
		return fmt.Errorf("unsupported retention unit")
	}

	i := strings.ReplaceAll(c.Retention, string(u), "")
	if _, err := strconv.ParseUint(i, 10, 32); err != nil {
		return fmt.Errorf("malformed retention field")
	}

	if c.Spool == "" {
		c.Spool = "/var/spool/peephole"
	}

	return nil
}
