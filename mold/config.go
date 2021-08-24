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
	i := strings.ReplaceAll(c.Retention, string(_util.Unit(c.Retention)), "")
	if _, err := strconv.ParseUint(i, 10, 32); err != nil {
		return fmt.Errorf("malformed retention field")
	}

	if c.Spool == "" {
		c.Spool = "/var/spool/peephole"
	}

	return nil
}
