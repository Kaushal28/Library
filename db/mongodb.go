package db

import (
	"os"
	"fmt"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Database = "library"
	DBPort = 27017
)

// NewMongoDBClient creates mongodb client object
func NewMongoDBClient() *mongo.Client {

	url := os.Getenv("MONGODB_HOSTNAME")
	if url == "" {
		panic("One of the required env. variable is empty: MONGODB_HOSTNAME")
	}
	url = fmt.Sprintf("mongodb://%s:%d", url, DBPort)

	// initializing db connector struct
	connector := newDBConnector(func(dbURI string) interface{} {
		// initiate mongodb connection
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbURI))
		if err != nil {
			// end the program if connection to db is not successful
			panic(err)
		}
		return client
	})
	// connecto to mongodb and extract mongo client from interface{}
	return connector.connect(url).(*mongo.Client)
}