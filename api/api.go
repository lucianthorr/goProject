package api

import (
	"github.com/lucianthorr/goProject/db"
)

type Config struct {
	// set dependencies
	dbCli db.Client
}

type Client interface {
	// set public methods
	Run() error
}

func New(cfg *Config) Client {
	c := &client{
		cfg: cfg,
	}
	// perform any other special bootstrapping
	return c
}

type client struct {
	cfg *Config
}

func (c *client) Run() error {
	return nil
}
