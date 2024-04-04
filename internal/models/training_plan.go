package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TrainingPlan struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Day           string
	Theme         string
	EstimatedTime string
	ScheduleDays  string
}
