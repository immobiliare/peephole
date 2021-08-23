package config

// Spy represents the substruct related to Salt events fetch API
type Spy struct {
	API    string `yaml:"api"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Client string `yaml:"client"`
}
