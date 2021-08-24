package spy

// Config represents the config struct for Spy module
type Config struct {
	API    string `yaml:"api"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Client string `yaml:"client"`
}

func (c *Config) Validate() error {
	return nil
}
