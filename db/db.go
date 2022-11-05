package db

import (
	"log"
)

type Config struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Client interface {
	// set public methods
	GetThing(condition string) (string, error)
}

func New(cfg *Config) (Client, error) {
	c := &client{
		cfg: cfg,
	}
	// perform any other special bootstrapping
	return c, nil
}

type client struct {
	cfg *Config
}

func (c *client) GetThing(condition string) (string, error) {
	log.Printf("Querying %s with condition %s", c.cfg.DBName, condition)
	return "successful result", nil
}
