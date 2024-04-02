package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/flavioesteves/wizer-dynamics-go/configs"
)

const version = "1.0.0"

type appConfig struct {
	port        int
	environment string
}

type application struct {
	config appConfig
	server *http.Server
}

func build() (*application, error) {

	serverConfig, err := config.GetConfiguration()
	if err != nil {
		fmt.Println("error:", err)
	}

	cfg := &appConfig{
		port:        serverConfig.Application.Port,
		environment: serverConfig.Environment,
	}

	app := &application{
		config: *cfg,
	}

	_, e := connectDB(&serverConfig.Database)
	if e != nil {
		fmt.Println("Failed to connect to the database")
	}

	appAddress := fmt.Sprintf("%s:%d", serverConfig.Application.Host, serverConfig.Application.Port)

	// Start server
	server := &http.Server{
		Addr:    appAddress,
		Handler: app.routes(),
	}
	fmt.Println(server.Addr)

	//TODO: refactor this logic
	app.server = server
	return app, nil
}

func Run() {
	app, err := build()
	if err != nil {
		fmt.Println("Error in build process")
	}
	app.server.ListenAndServe()
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
