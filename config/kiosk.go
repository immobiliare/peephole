package config

// Kiosk represents the substruct related to the webserver exposed
type Kiosk struct {
	Bind      string            `yaml:"bind"`
	BasicAuth map[string]string `yaml:"basic_auth"`
}
