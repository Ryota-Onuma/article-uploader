package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server *Server
}

type Server struct {
	Port int    `envconfig:"PORT" default:"3000"`
	Env  string `envconfig:"ENV" default:"local"`
}

func (server *Server) IsLocal() bool {
	return server.Env == "local"
}

func New() *Config {
	return &Config{}
}

func (config *Config) Setup() error {
	if err := envconfig.Process("", config); err != nil {
		return err
	}

	return nil
}
