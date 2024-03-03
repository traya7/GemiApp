package domain

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB represents the MongoDB database connection.
type MongoDB struct {
	client *mongo.Client
	dbName string
}

// NewMongoDB creates a new MongoDB instance and connects to the database.
func NewMongoDB(connectionString, dbName string) (*MongoDB, error) {
	client, err := connectMongoDB(connectionString)
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		client: client,
		dbName: dbName,
	}, nil
}

// connectMongoDB establishes a connection to the MongoDB database.
func connectMongoDB(connectionString string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, err
	}
	return client, nil
}

// GetCollection returns a MongoDB collection based on the specified database and collection names.
func (db *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return db.client.Database(db.dbName).Collection(collectionName)
}

// Close closes the MongoDB client connection.
func (db *MongoDB) Close() error {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
	}
	return err
}
