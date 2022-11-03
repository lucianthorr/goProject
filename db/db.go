package db

type Config struct {
	// set dependencies

}

type Client interface {
	// set public methods
	Query(q string) error
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

func (c *client) Query(q string) error {
	return nil
}
