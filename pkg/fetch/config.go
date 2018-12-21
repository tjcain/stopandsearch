package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// // LoadConfig attempts to load .config
// func LoadConfig() (*Config, error) {
// 	f, err := os.Open(".config")
// 	if err != nil {
// 		fmt.Println("Using the default configs...")
// 		return DefaultConfig(), nil
// 	}

// 	var c Config
// 	d := json.NewDecoder(f)
// 	if err := d.Decode(&c); err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("Successfully loaded .config")
// 	return &c, nil
// }

// // Config holds start up configuration variables
// type Config struct {
// 	Env      string         `json:"env"`
// 	Database PostgresConfig `json:"database"`
// 	// Mailgun  MailgunConfig  `json:"mailgun"`
// }

// // DefaultConfig generates development environment variables
// func DefaultConfig() *Config {
// 	return &Config{
// 		Env:      "dev",
// 		Database: DefaultPostgresConfig(),
// 	}
// }

// // PostgresConfig holds config variables for connecting to postgres
// type PostgresConfig struct {
// 	Host     string `json:"host"`
// 	Port     int    `json:"port"`
// 	User     string `json:"user"`
// 	Password string `json:"password"`
// 	Database string `json:"database"`
// }

// // DefaultPostgresConfig returns a PostgresConfig with default values
// func DefaultPostgresConfig() PostgresConfig {
// 	return PostgresConfig{
// 		Host:     "localhost",
// 		Port:     5432,
// 		User:     "postgres",
// 		Password: "",
// 		Database: "stopandsearch_test",
// 	}
// }
