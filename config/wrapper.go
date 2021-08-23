package config

// Wrapper represents the abstraction of the parsed
// configuration file
type Wrapper struct {
	Spy   []*Spy `yaml:"spy"`
	Kiosk *Kiosk `yaml:"kiosk"`
	Mold  *Mold  `yaml:"mold"`
}
