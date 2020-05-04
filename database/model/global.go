package model

import (
	"github.com/qq15725/go/database"
	"github.com/qq15725/go/database/dsn"
)

var (
	cfg         = newConfig()
	connections = make(map[string]*database.Connection)
)

func SetConnections(connections map[string]dsn.Connector) {
	cfg.SetConnections(connections)
}

func CloseConnections() {
	for name := range connections {
		connections[name].Close()
	}
}
