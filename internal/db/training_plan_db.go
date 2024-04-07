package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
)

func (s *MongoDBStorer) InsertTraining(ctx context.Context, tp *models.TrainingPlan) error {
	res, err := s.DB.Collection(s.Coll).InsertOne(ctx, tp)
	if err != nil {
		return err
	}

	tp.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

func (s *MongoDBStorer) GetALlTrainings(ctx context.Context) ([]*models.TrainingPlan, error) {
	cursor, err := s.DB.Collection(s.Coll).Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}

	trainingPlans := []*models.TrainingPlan{}
	err = cursor.All(ctx, &trainingPlans)
	return trainingPlans, err
}

func (s *MongoDBStorer) GetTrainigByID(ctx context.Context, id string) (*models.TrainingPlan, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res      = s.DB.Collection(s.Coll).FindOne(ctx, bson.M{"_id": objID})
		tp       = &models.TrainingPlan{}
		err      = res.Decode(tp)
	)

	return tp, err
}

func (s *MongoDBStorer) DeleteTrainingPlanByID(ctx context.Context, id string) (*mongo.DeleteResult, error) {

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

	return res, err
}
