package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	_Id          primitive.ObjectID `bson:"_id,omitempty"`
	Guid         string
	RefreshToken string
}
