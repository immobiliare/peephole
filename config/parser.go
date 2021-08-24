package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Parse generates a new Config instance
// starting from a configuration file path
func Parse(path string) (*Wrapper, error) {
	config := new(Wrapper)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return process(config)
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
