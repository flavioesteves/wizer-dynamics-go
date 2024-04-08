package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TrainingSession struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Day           string             `json:"day" bson:"day"`
	Theme         string             `json:"theme" bson:"theme"`
	EstimatedTime string             `json:"estimated_time" bson:"estimated_time"`
	ScheduleDays  string             `json:"schedule_days" bson:"schedule_days"`
}
