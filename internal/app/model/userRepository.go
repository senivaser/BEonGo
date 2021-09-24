package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUser() (*UserRepository, error) {

	config := NewConfig()
	db := NewDB()

	collection, err := db.GetCollection(config, "User")

	if err != nil {
		return nil, err
	}

	return &UserRepository{
		collection: collection,
	}, nil
}

func (ur *UserRepository) Get(guid string) (*User, error) {
	var user *User
	filter := bson.M{"guid": guid}
	err := ur.collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return nil, err
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
