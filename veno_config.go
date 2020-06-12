package veno

import (
	"github.com/qq15725/veno/config"
)

type Config struct {
	*config.Config
}

func newConfig() *Config {
	return &Config{
		Config: config.New(),
	}
}
