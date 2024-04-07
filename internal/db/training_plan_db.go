package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
)

func (s *MongoDBStorer) Insert(ctx context.Context, tp *models.TrainingPlan) error {
	res, err := s.DB.Collection(s.Coll).InsertOne(ctx, tp)
	if err != nil {
		return err
	}

	tp.ID = res.InsertedID.(primitive.ObjectID)
	return err
}
