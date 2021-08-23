package config

// Mold represents the substruct related to the db
type Mold struct {
	Spool     string `yaml:"spool"`
	Retention string `yaml:"retention"`
}
