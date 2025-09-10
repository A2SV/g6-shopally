package mongodb // Package name should be 'mongodb' as per folder structure

import (
	"context"
	"fmt"
	"log" // Using standard log for any potential errors/info as per repository's current state
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref" // Added for pinging primary
)

// Connect establishes a new MongoDB client connection.
// It also ensures that a unique index is created on the 'user_email' field
// in the specified collection upon successful connection.
func Connect(uri, dbName, collectionName string) (*mongo.Client, error) { // Added dbName, collectionName
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	// Ping the primary to verify connection status
	if err := client.Ping(ctx, readpref.Primary()); err != nil { // Used readpref.Primary()
		// Disconnect if ping fails
		if disconnectErr := client.Disconnect(context.Background()); disconnectErr != nil {
			log.Printf("Warning: Failed to disconnect MongoDB client after ping failure: %v", disconnectErr)
		}
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	log.Printf("MongoDB client connected successfully to %s.", uri)

	collection := client.Database(dbName).Collection(collectionName)

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "user_email", Value: 1}}, 
		Options: options.Index().SetUnique(true),       
	}

	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Printf("Warning: Error creating unique index on 'user_email' in collection '%s': %v. This might be normal if the index already exists.", collectionName, err)
	} else {
		log.Printf("Unique index created on 'user_email' in collection '%s'.", collectionName)
	}


	return client, nil
}

// Disconnect closes the MongoDB client connection gracefully.
func Disconnect(client *mongo.Client) error {
	if client == nil {
		log.Println("Attempted to disconnect a nil MongoDB client.")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Printf("Error disconnecting MongoDB client: %v", err)
		return fmt.Errorf("failed to disconnect MongoDB client: %w", err)
	}
	log.Println("MongoDB client disconnected successfully.")
	return nil
}