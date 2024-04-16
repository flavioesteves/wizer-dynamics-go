package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	config "github.com/flavioesteves/wizer-dynamics-go/configs"
)

type MongoDBStorer struct {
	DB          *mongo.Database
	Coll        string
	RedisClient *redis.Client
}

func NewMongoDBStore(db *mongo.Database, redisClient *redis.Client, coll string) *MongoDBStorer {
	return &MongoDBStorer{
		DB:          db,
		RedisClient: redisClient,
		Coll:        coll,
	}
}

func ConnectDB(dbSettings *config.DatabaseSettings) (*mongo.Client, error) {
	mongoURI := fmt.Sprintf("%s://%s:%d", dbSettings.Model, dbSettings.Host, dbSettings.Port)

	// Try to connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return client, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		fmt.Printf("Can't connect to db: %v", err)
		return client, err
	}

	fmt.Println("Connected to MongoDB at: ", mongoURI)
	return client, nil

}
