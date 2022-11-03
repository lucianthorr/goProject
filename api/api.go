package api

import (
	"log"

	"github.com/lucianthorr/goProject/db"
)

type Config struct {
	Port int `yaml:"port"`
}

type Client interface {
	// set public methods
	Run() error
}

func New(cfg *Config, dbCli db.Client) Client {
	c := &client{
		cfg:   cfg,
		dbCli: dbCli,
	}
	// perform any other special bootstrapping
	return c
}

type client struct {
	cfg   *Config
	dbCli db.Client
}

func (c *client) Run() error {
	log.Printf("Here's your port number: %d", c.cfg.Port)
	return c.dbCli.Query("SELECT * from orders")
}
