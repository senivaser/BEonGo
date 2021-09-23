package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	_Id          primitive.ObjectID `bson:"_id,omitempty"`
	guid         string
	refreshToken string
}
