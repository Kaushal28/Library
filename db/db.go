package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri = "mongodb://localhost:27017"
)

// creating a thread safe singleton type for mongo client
type db struct {
	instance *mongo.Client
}

var dbInstance *db = nil
var once sync.Once

const (
	Database = "library"
)

func New() *mongo.Client {
	if dbInstance == nil {
		once.Do(func() {
			client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
			if err != nil {
				// end the program if connection to db is not successful
				panic(err)
			}
			dbInstance = &db{client}
		})
	}
	return dbInstance.instance
}
