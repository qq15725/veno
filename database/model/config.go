package model

import (
	"fmt"
	"github.com/qq15725/go/database/dsn"
)

type config struct {
	connections map[string]dsn.Connector
}

func (cfg *config) SetConnections(connections map[string]dsn.Connector) {
	cfg.connections = connections
}

func (cfg *config) SetConnection(name string, driver dsn.Connector) {
	cfg.connections[name] = driver
}

func (cfg *config) GetConnection(name string) (dsn.Connector, error) {
	driver, ok := cfg.connections[name]
	if !ok {
		return nil, fmt.Errorf("database: unknown connection %q (forgotten set?)", name)
	}
	return driver, nil
}

func newConfig() *config {
	return &config{connections: make(map[string]dsn.Connector)}
}
