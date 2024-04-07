package db

import (
	"context"

	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *MongoDBStorer) InsertExercise(ctx context.Context, e *models.Exercise) (*models.Exercise, error) {
	res, err := s.DB.Collection(s.Coll).InsertOne(ctx, e)
	if err != nil {
		return nil, err
	}

	e.ID = res.InsertedID.(primitive.ObjectID)

	return e, err
}

func (s *MongoDBStorer) GetALlExercises(ctx context.Context) ([]*models.Exercise, error) {
	cursor, err := s.DB.Collection(s.Coll).Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}

	exercises := []*models.Exercise{}
	err = cursor.All(ctx, &exercises)
	return exercises, err
}

func (s *MongoDBStorer) GetExerciseByID(ctx context.Context, id string) (*models.Exercise, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res      = s.DB.Collection(s.Coll).FindOne(ctx, bson.M{"_id": objID})
		e        = &models.Exercise{}
		err      = res.Decode(e)
	)
	return e, err
}

func (s *MongoDBStorer) DeleteExerciseByID(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res, err = s.DB.Collection(s.Coll).DeleteOne(ctx, bson.M{"_id": objID})
	)

	if err != nil {
		return res, err
	}

	if res.DeletedCount == 0 {
		return res, err
	}

	// TODO: Improve response currently is returning
	// {"DeletedCount":1}
	return res, err
}
