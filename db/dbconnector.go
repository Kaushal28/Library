package db

import (
	"sync"
)

// creating a thread safe singleton type for mongo client
type dbClient struct {
	 client interface{}
}
var clientInstance *dbClient = nil
var once sync.Once

type connectorFunc func(dbURI string) interface{}
type dbConnector struct {
	// db connector method (of type connect)
	connector connectorFunc
}

func newDBConnector(c connectorFunc) *dbConnector {
	return &dbConnector{connector: c}
}

func (d *dbConnector) connect(dbURI string) interface{} {
	if clientInstance == nil {
		once.Do(func() {
			client := d.connector(dbURI)
			clientInstance = &dbClient{client: client}
		})
	}
	return clientInstance.client
}
