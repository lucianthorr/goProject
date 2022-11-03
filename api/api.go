package api

import (
	"github.com/lucianthorr/goProject/db"
)

type Config struct {
	// set dependencies
	DB db.Client
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
	return nil
}
