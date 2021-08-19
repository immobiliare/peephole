package config

// Wrapper represents the abstraction of the parsed
// configuration file
type Wrapper struct {
	Spy   []*Spy `yaml:"spy"`
	Kiosk *Kiosk `yaml:"kiosk"`
}

// Spy represents the substruct related to Salt events fetch API
type Spy struct {
	API    string `yaml:"api"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Client string `yaml:"client"`
}

// Kiosk represents the substruct related to the webserver exposed
type Kiosk struct {
	Bind      string            `yaml:"bind"`
	BasicAuth map[string]string `yaml:"basic_auth"`
}
