package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
)

func (s *MongoDBStorer) InsertTraining(ctx context.Context, ts *models.TrainingSession) (*models.TrainingSession, error) {
	res, err := s.DB.Collection(s.Coll).InsertOne(ctx, ts)
	if err != nil {
		return nil, err
	}

	ts.ID = res.InsertedID.(primitive.ObjectID)
	return ts, err
}

func (s *MongoDBStorer) GetALlTrainings(ctx context.Context) ([]*models.TrainingSession, error) {
	cursor, err := s.DB.Collection(s.Coll).Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}

	trainingSessions := []*models.TrainingSession{}
	err = cursor.All(ctx, &trainingSessions)
	return trainingSessions, err
}

func (s *MongoDBStorer) GetTrainigByID(ctx context.Context, id string) (*models.TrainingSession, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	res := s.DB.Collection(s.Coll).FindOne(ctx, bson.M{"_id": objID})
	ts := &models.TrainingSession{}
	err := res.Decode(ts)

	return ts, err
}

func (s *MongoDBStorer) DeleteTrainingByID(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	res, err := s.DB.Collection(s.Coll).DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil {
		return res, err
	}

	if res.DeletedCount == 0 {
		return res, err
	}

	// TODO: Improve response, currently is returning {"DeleteCount: x"}
	return res, err
}
func (s *MongoDBStorer) UpdateTrainingByID(ctx context.Context, ts *models.TrainingSession) (*mongo.UpdateResult, error) {

	update := bson.M{"$set": bson.M{
		"day":            ts.Day,
		"theme":          ts.Theme,
		"estimated_time": ts.EstimatedTime,
		"schedule_days":  ts.ScheduleDays,
	}}

	res, err := s.DB.Collection(s.Coll).UpdateOne(ctx, bson.M{"_id": ts.ID}, update)
	if err != nil {
		return nil, err
	}

	return res, err
}
