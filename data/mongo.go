package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial-part-1-connecting-using-bson-and-crud-operations
var (
	connectionString = "mongodb://root:secret@mongodb:27017/my-mongo-db1"
	Client           *MongoDBClient
)

func init() {
	client, err := NewMongoDBClient(connectionString)
	if err != nil {
		log.Fatalf("Failed to init mongo db client: %s", err.Error())
	}
	Client = client
}

type MongoDBClient struct {
	client *mongo.Client
}

type MongoDBClientInterface interface {
	Disconnect() error
	Ping() error
}

func NewMongoDBClient(connectionString string) (*MongoDBClient, error) {

	mongoDBClient := &MongoDBClient{}
	err := mongoDBClient.connect(connectionString)
	if err != nil {
		return nil, err
	}

	err = mongoDBClient.Ping()
	if err != nil {
		return nil, err
	}

	return mongoDBClient, nil
}

func (mongoDBClient *MongoDBClient) GetInternalClient() *mongo.Client {
	return mongoDBClient.client
}

func (mongoDBClient *MongoDBClient) connect(connectionString string) error {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	mongoDBClient.client = client
	return nil
}

func (mongoDBClient *MongoDBClient) Ping() error {
	// Check the connection
	err := mongoDBClient.client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	return nil
}

func (mongoDBClient *MongoDBClient) Disconnect() error {
	err := mongoDBClient.client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
