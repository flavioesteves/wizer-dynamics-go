package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Steps string             `json:"steps" bson:"steps"`
	Video string             `json:"video" bson:"video"`
	Photo string             `json:"photo" bson:"photo"`
}
