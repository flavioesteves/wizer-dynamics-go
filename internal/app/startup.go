package app

import (
	"context"
	"fmt"
	"github.com/flavioesteves/wizer-dynamics-go/config"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Application struct {
	Server *http.Server
	Port   int
}

func Build(c *config.Settings) (*Application, error) {
	dbClient, err := getDbClient(&c.Database)
	if err != nil {
		fmt.Println("Failed to connect to the database")
	}

	fmt.Println(dbClient)

	appAddress := fmt.Sprintf("%s:%d", c.Application.Host, c.Application.Port)

	fmt.Println(appAddress)
	router := gin.Default()

	// Start the server
	server := &http.Server{
		Addr:    appAddress,
		Handler: router,
	}

	return &Application{Server: server, Port: c.Application.Port}, nil
}

// Connect to database
func getDbClient(dbSettings *config.DatabaseSettings) (*mongo.Client, error) {
	mongoURI := fmt.Sprintf("%s://%s:%d", dbSettings.Model, dbSettings.Host, dbSettings.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))

	if err != nil {
		fmt.Println("Failed to connect to the database")
	}
	return client, nil
}
