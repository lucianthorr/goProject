package configs

import (
	"io/ioutil"

	"github.com/lucianthorr/goProject/api"
	"github.com/lucianthorr/goProject/db"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	API *api.Config `yaml:"api"`
	DB  *db.Config  `yaml:"db"`
}

func Read(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading config")
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling config")
	}
	return cfg, nil
}
