package db

import (
	"context"
	"fmt"
	"os"

	"github.com/Kaushal28/library/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Database = "library"
	DBPort   = 27017  // this should be configurable as well
)

// newMongoDBClient creates mongodb client object
func newMongoDBClient() *mongo.Client {

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

// mongodb data access layer. It uses common data access layer definition from db.go
type MongoDAL struct {
	client *mongo.Client
	dbName string
}

// NewMongoDAL creates a MongoDAL instance
func NewMongoDAL() DataAccessLayer {
	mongo := &MongoDAL{
		client: newMongoDBClient(),
		dbName: Database,
	}
	return mongo
}

// implementing data access layer's methods on MongoDAL to satisfy DataAccessLayer interface

// coll is a helper method to get a collection object
func (m *MongoDAL) coll(collection string) *mongo.Collection {
	return m.client.Database(m.dbName).Collection(collection)
}

// Findbooks queries mongodb with given query to fetch relevant books
func (m *MongoDAL) FindBooks(coll string, ctx context.Context, query interface{}) ([]entities.Book, error) {
	cursor, err := m.coll(coll).Find(ctx, query)
	if err != nil {
		return nil, err
	}
	var results []entities.Book
	err = cursor.All(context.Background(), &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// InsertBooks inserts given list of books into mongodb
func (m *MongoDAL) InsertBooks(coll string, ctx context.Context, documents []interface{}) (interface{}, error) {

	_, err := m.coll(coll).InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}
	// first return can be id of the inserted book or similar data.
	return nil, nil
}
