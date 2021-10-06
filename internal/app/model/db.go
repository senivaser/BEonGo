package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	config *Config
	client *mongo.Client
}

func NewDB() *DB {
	return &DB{
		config: NewConfig(),
	}
}

func (db *DB) getClient(uri string) (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	db.client = client

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = db.client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client, err
}

func (db *DB) GetCollection(config *Config, collectionName string) (*mongo.Collection, error) {

	client, clientErr := db.getClient(config.Uri)
	var collection *mongo.Collection

	if clientErr == nil {
		collection = client.Database(config.Database).Collection(collectionName)
	} else {
		collection = nil
	}
	return collection, clientErr
}
