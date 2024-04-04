package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string
	Password string
	Updated  string // TODO Implement time
	Created  string // TODO Implement time
}
