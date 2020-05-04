package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	items map[string]interface{}
}

func (cfg *Config) Load(key string, path string) {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	var value interface{}
	decoder.Decode(&value)
	cfg.Set(key, value)
}

func (cfg *Config) Set(key string, value interface{}) {
	cfg.items[key] = value
}

func (cfg *Config) Get(key string) interface{} {
	return cfg.items[key]
}

func New() *Config {
	return &Config{
		items: make(map[string]interface{}),
	}
}
