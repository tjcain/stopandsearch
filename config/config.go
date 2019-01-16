package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config represents a stopandsearch configuration
type Config struct {

	// DB is the database configuration
	DB struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
	} `json:"db"`
}

// ParseFile reads a json config file and returns a Configuration
func ParseFile(filepath string) (*Config, error) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println("bad file", err)
		return nil, err
	}
	d := json.NewDecoder(f)
	c := &Config{}
	if err = d.Decode(&c); err != nil {
		return nil, err
	}
	return c, nil
}
