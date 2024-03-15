package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	//"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/flavioesteves/wizer-dynamics-go/configs"
	"github.com/flavioesteves/wizer-dynamics-go/internal/routes"
)

type Application struct {
	Server *http.Server
	Port   int
}

func Build(c *config.Settings) (*Application, error) {
	//TODO use this connection for models/controllers
	_, err := connectDB(&c.Database)
	if err != nil {
		fmt.Println("Failed to connect to the database")
	}
	appAddress := fmt.Sprintf("%s:%d", c.Application.Host, c.Application.Port)

	router := gin.Default()
	exerciseGroup := router.Group("/exercise")
	routes.RegisterExerciseRoutes(exerciseGroup)

	// Start the server
	server := &http.Server{
		Addr:    appAddress,
		Handler: router,
	}

	return &Application{Server: server, Port: c.Application.Port}, nil
}

func connectDB(dbSettings *config.DatabaseSettings) (*mongo.Client, error) {
	mongoURI := fmt.Sprintf("%s://%s:%d", dbSettings.Model, dbSettings.Host, dbSettings.Port)

	// Try to connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		fmt.Println("Failed to connect to the database")
	}

	fmt.Println("Connected to MongoDB at: ", mongoURI)
	return client, nil
}
