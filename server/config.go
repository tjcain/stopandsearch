package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	// ErrInvalidConfigFile is returned when a configuration file is invalid.
	ErrInvalidConfigFile = errors.New("no configuration file specified")
	// DefaultConfig is initialized upon startup of the program and used
	// when a config filepath is not given.
	DefaultConfig = Config{}
)

// Config holds the configuration options for the application.
type Config struct {
	Bind string `json:"bind,omitempty"`

	LogLevel string `json:"log_level,omitempty"`
	LogFile  string `json:"log_file"`

	CacheMaxAge int `json:"cache_max_age,omitempty"`

	TmpDir string `json:"tmp_dir,omitempty"`

	Limits struct {
		// any limits will go here
	} `json:"limits,omitempty"`

	PgDB struct {
		Host     string `json:"host,omitempty"`
		Port     int    `json:"port,omitempty"`
		User     string `json:"user,omitempty"`
		Password string `json:"password,omitempty"`
		Database string `json:"database,omitempty"`
	} `json:"pg_db,omitempty"`

	Prometheus struct {
		//TODO: Add prometheus configuration here
	} `json:"prometheus_d,omitempty"`

	// Mail
	// Boltdb
}

func init() {
	cf := Config{
		Bind:        "0.0.0.0:4000",
		LogLevel:    "INFO",
		LogFile:     "",
		CacheMaxAge: 0,
		TmpDir:      "",
	}

	cf.PgDB.Host = "localhost"
	cf.PgDB.Port = 5432
	cf.PgDB.Password = ""
	cf.PgDB.User = "postgres"
	cf.PgDB.Database = "stopandsearch_test"

	DefaultConfig = cf
}

// NewConfig loads a New configuration using initialized default values.
func NewConfig() *Config {
	cf := DefaultConfig
	return &cf
}

// NewConfigFromFile loads a new configuration from a JSON file
func NewConfigFromFile(confFile string) (*Config, error) {
	f, err := os.Open(confFile)
	if err != nil {
		return nil, err
	}

	c := NewConfig()

	d := json.NewDecoder(f)
	if err := d.Decode(&c); err != nil {
		return nil, err
	}

	return c, nil

}

// GetDB returns a new database
func (cf *Config) GetDB() (*DB, error) {
	s := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cf.PgDB.User, cf.PgDB.Password, cf.PgDB.Database,
		cf.PgDB.Host, cf.PgDB.Port)
	db, err := NewPostgresDB(s)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
