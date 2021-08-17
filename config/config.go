package config

// Wrapper represents the abstraction of the parsed
// configuration file
type Wrapper struct {
	Spy []*Spy `yaml:"spy"`
}

// Spy represents the substruct related Salt events fetch API
type Spy struct {
	API    string `yaml:"api"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Client string `yaml:"client"`
}
