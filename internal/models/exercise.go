package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Steps string             `bson:"steps"`
	Video string             `bson:"video"`
	Photo string             `bson:"photo"`
}
