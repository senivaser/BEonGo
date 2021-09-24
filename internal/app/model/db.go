package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	config *Config
}

func NewDB() *DB {
	return &DB{
		config: NewConfig(),
	}
}

func (db *DB) getClient(uri string) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client, err
}

func (db *DB) GetCollection(config *Config, collectionName string) (*mongo.Collection, error) {
	client, clientErr := db.getClient(config.uri)
	var collection *mongo.Collection
	if clientErr != nil {
		collection = client.Database(config.database).Collection(collectionName)
	} else {
		collection = nil
	}

	return collection, clientErr
}
