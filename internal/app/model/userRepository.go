package model

import (
	"context"
	"fmt"
	"time"

	"github.com/go-errors/errors"
	"github.com/senivaser/BEonGo/internal/app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	collection *mongo.Collection
	Config     *Config
	DB         *DB
}

func NewUser(config *Config) (*UserRepository, error) {

	db := NewDB()
	collection, err := db.GetCollection(config, "users")

	if err != nil {
		return nil, err
	}

	return &UserRepository{
		collection: collection,
	}, nil
}

func (ur *UserRepository) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (ur *UserRepository) GetBy(field string, value string) (User, *errors.Error) {

	filter := bson.D{{field, value}}
	user, err := ur.Get(filter)

	return user, err
}

func (ur *UserRepository) Get(filter primitive.D) (User, *errors.Error) {
	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("filter: " + utils.ToString(filter))
	err := ur.collection.FindOne(ctx, filter).Decode(&user)
	sErr := errors.Wrap(err, 0)

	if err != nil {
		return User{}, sErr
	}

	return user, nil
}

func (ur *UserRepository) UpdateBy(field string, value string, updateField string, updateValue string) (mongo.UpdateResult, error) {

	var err error

	if updateField == "refreshToken" {
		updateValue, err = ur.hashPassword(updateValue)
	}

	filter := bson.D{{field, value}}
	setter := bson.D{{updateField, updateValue}}
	result, err := ur.Update(filter, setter)

	return result, err
}

func (ur *UserRepository) Update(filter primitive.D, setter primitive.D) (mongo.UpdateResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := ur.collection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", setter},
		},
	)

	if err != nil {
		return *result, err
	}

	return *result, nil
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
