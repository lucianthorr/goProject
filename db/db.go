package db

import (
	"errors"
	"log"
	"os"
)

type Config struct {
	Hostname string `yaml:"hostname"`
	DBName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `yaml:"-"`
}

type Client interface {
	// set public methods
	Query(q string) error
}

func New(cfg *Config) (Client, error) {
	c := &client{
		cfg: cfg,
	}
	cfg.Password = os.Getenv("DB_PASSWORD")
	if cfg.Password == "" {
		return nil, errors.New("Missing DB_PASSWORD environment variable.")
	}
	// perform any other special bootstrapping
	return c, nil
}

type client struct {
	cfg *Config
}

func (c *client) Query(q string) error {
	log.Printf("Querying %s against %s", q, c.cfg.DBName)
	return nil
}
