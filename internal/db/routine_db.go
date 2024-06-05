package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
)

func (s *MongoDBStorer) InsertRoutine(ctx context.Context, r *models.Routine) (*models.Routine, error) {
	res, err := s.DB.Collection(s.Coll).InsertOne(ctx, r)
	if err != nil {
		return nil, err
	}

	r.ID = res.InsertedID.(primitive.ObjectID)
	return r, err
}

func (s *MongoDBStorer) GetALlRoutines(ctx context.Context) ([]*models.Routine, error) {
	cursor, err := s.DB.Collection(s.Coll).Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}

	routines := []*models.Routine{}
	err = cursor.All(ctx, &routines)
	return routines, err
}

func (s *MongoDBStorer) GetRoutineByID(ctx context.Context, id string) (*models.Routine, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	res := s.DB.Collection(s.Coll).FindOne(ctx, bson.M{"_id": objID})
	r := &models.Routine{}
	err := res.Decode(r)

	return r, err
}

func (s *MongoDBStorer) DeleteRoutineByID(ctx context.Context, id string) (*mongo.DeleteResult, error) {
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
func (s *MongoDBStorer) UpdateRoutineByID(ctx context.Context, r *models.Routine) (*mongo.UpdateResult, error) {

	update := bson.M{"$set": bson.M{
		"day":            r.Day,
		"theme":          r.Theme,
		"estimated_time": r.EstimatedTime,
		"schedule_days":  r.ScheduleDays,
	}}

	res, err := s.DB.Collection(s.Coll).UpdateOne(ctx, bson.M{"_id": r.ID}, update)
	if err != nil {
		return nil, err
	}

	return res, err
}
