package config

import (
	_kiosk "github.com/streambinder/peephole/kiosk"
	_mold "github.com/streambinder/peephole/mold"
	_spy "github.com/streambinder/peephole/spy"
)

// Wrapper represents the abstraction of the parsed
// configuration file
type Wrapper struct {
	Spy   []*_spy.Config `yaml:"spy"`
	Kiosk *_kiosk.Config `yaml:"kiosk"`
	Mold  *_mold.Config  `yaml:"mold"`
}
