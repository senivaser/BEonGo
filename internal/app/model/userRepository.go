package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	Config     *Config
	DB         *DB
}

func NewUser(config *Config) (*UserRepository, error) {

	db := NewDB()
	fmt.Print("config: ", config)
	collection, err := db.GetCollection(config, "User")

	if err != nil {
		return nil, err
	}

	return &UserRepository{
		collection: collection,
		Config:     config,
		DB:         db,
	}, nil
}

func (ur *UserRepository) Get(guid string) (User, error) {
	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{"guid", guid}}
	err := ur.collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (ur *UserRepository) Create(user *User) (*mongo.InsertOneResult, error) {

	result, err := ur.collection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// func (ur *UserRepository) Get1(guid string) (User, error) {
// 	var user User
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	defer func() {
// 		if err = client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()
// 	err = client.Ping(ctx, readpref.Primary())

// 	collection := client.Database("taskDB").Collection("User")
// 	fmt.Println("collection1: ", collection)
// 	filter := bson.D{{"guid", guid}}
// 	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	err = collection.FindOne(ctx, filter).Decode(&user)
// 	if err == mongo.ErrNoDocuments {
// 		// Do something when no record was found
// 		fmt.Println("record does not exist")
// 	} else if err != nil {
// 		log.Fatal(err)
// 	}

// 	return user, nil

// }
