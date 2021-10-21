package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
			client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", os.Getenv("MONGODB_HOSTNAME"))))
			if err != nil {
				// end the program if connection to db is not successful
				panic(err)
			}
			dbInstance = &db{client}
		})
	}
	return dbInstance.instance
}
