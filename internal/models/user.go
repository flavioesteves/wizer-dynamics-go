package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email" unique:"true"`
	Password  string             `json:"password" bson:"password"`
	Updated   string             `json:"updated" bson:"updated"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
}
