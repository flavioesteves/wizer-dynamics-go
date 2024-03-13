package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ID primitive.ObjectID `bson:"_id"`
}
