package db

import (
	"context"
	"sync"

	"github.com/library/entities"
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

// Defining data access layer. This will be used instead of directly accessing
// database clients. This makes code more decoupled and testable.
// Reference: https://medium.com/@harrygogonis/testing-go-mocking-third-party-dependancies-4ab4e1c9bd3f
type DataAccessLayer interface {
	// declare all data access methods used throughout the routers
	FindBooks(coll string, ctx context.Context, query interface{}) ([]entities.Book, error)
	InsertBooks(coll string, ctx context.Context, documents []interface{}) (interface{}, error)
}
