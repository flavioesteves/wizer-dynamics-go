package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
)

func (s *MongoDBStorer) InsertUser(ctx context.Context, u *models.User) (*models.User, error) {

	filter := bson.M{"email": u.Email}
	existingUser := &models.User{}

	err := s.DB.Collection(s.Coll).FindOne(ctx, filter).Decode(existingUser)

	fmt.Println(err)
	if err == nil {
		fmt.Println("Inside Nil")
		return nil, errors.New("email already exists")
	}

	res, err := s.DB.Collection(s.Coll).InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}

	u.ID = res.InsertedID.(primitive.ObjectID)

	return u, err
}

func (s *MongoDBStorer) GetALlUsers(ctx context.Context) ([]*models.User, error) {
	cursor, err := s.DB.Collection(s.Coll).Find(ctx, map[string]any{})
	if err != nil {

		return nil, err
	}

	users := []*models.User{}
	err = cursor.All(ctx, &users)
	return users, err
}

func (s *MongoDBStorer) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	res := s.DB.Collection(s.Coll).FindOne(ctx, bson.M{"_id": objID})
	e := &models.User{}
	err := res.Decode(e)
	return e, err
}

func (s *MongoDBStorer) UpdateUserByID(ctx context.Context, u *models.User) (*mongo.UpdateResult, error) {
	update := bson.M{"$set": bson.M{
		"email":      u.Email,
		"password":   u.Password,
		"updated":    u.Updated,
		"created_at": u.CreatedAt,
	}}

	res, err := s.DB.Collection(s.Coll).UpdateOne(ctx, bson.M{"_id": u.ID}, update)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *MongoDBStorer) IsValidCredentials(ctx context.Context, email string, password string) bool {
	cur := s.DB.Collection(s.Coll).FindOne(ctx, bson.M{
		"email":    email,
		"password": password,
	})

	return cur.Err() == nil
}
