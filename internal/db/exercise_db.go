package db

import (
	"context"

	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MongoDBStorer) Insert(ctx context.Context, e *models.Exercise) error {
	res, err := s.DB.Collection(s.Coll).InsertOne(ctx, e)
	if err != nil {
		return err
	}

	e.ID = res.InsertedID.(primitive.ObjectID)

	return err
}

func (s *MongoDBStorer) GetALl(ctx context.Context) ([]*models.Exercise, error) {
	cursor, err := s.DB.Collection(s.Coll).Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}

	exercises := []*models.Exercise{}
	err = cursor.All(ctx, &exercises)
	return exercises, err
}

func (s *MongoDBStorer) GetByID(ctx context.Context, id string) (*models.Exercise, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res      = s.DB.Collection(s.Coll).FindOne(ctx, bson.M{"_id": objID})
		e        = &models.Exercise{}
		err      = res.Decode(e)
	)
	return e, err
}
