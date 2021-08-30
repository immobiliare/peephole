package config

import (
	"io/ioutil"
	"os"

	_kiosk "github.com/streambinder/peephole/kiosk"
	_mold "github.com/streambinder/peephole/mold"
	_spy "github.com/streambinder/peephole/spy"
	"gopkg.in/yaml.v2"
)

// Wrapper represents the abstraction of the parsed
// configuration file
type Wrapper struct {
	Debug bool           `yaml:"debug"`
	Spy   []*_spy.Config `yaml:"spy"`
	Kiosk *_kiosk.Config `yaml:"kiosk"`
	Mold  *_mold.Config  `yaml:"mold"`
}

// Parse generates a new Config instance
// starting from a configuration file path
func Parse(path string) (*Wrapper, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parseYaml(content)
}

func parseYaml(content []byte) (*Wrapper, error) {
	config := Wrapper{
		Spy:   []*_spy.Config{},
		Mold:  new(_mold.Config),
		Kiosk: new(_kiosk.Config),
	}
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}

	return process(&config)
}

func process(cfg *Wrapper) (*Wrapper, error) {
	for _, e := range cfg.Spy {
		if err := e.Validate(); err != nil {
			return nil, err
		}
	}
	if err := cfg.Kiosk.Validate(); err != nil {
		return nil, err
	}
	if err := cfg.Mold.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}
