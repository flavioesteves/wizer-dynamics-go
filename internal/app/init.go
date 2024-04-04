package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/flavioesteves/wizer-dynamics-go/configs"
	"github.com/flavioesteves/wizer-dynamics-go/internal/routers"
)

type appConfig struct {
	port        int
	environment string
	Settings    config.Settings
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
		Settings:    serverConfig,
	}

	dbClient, e := connectDB(&serverConfig.Database)
	if e != nil {
		fmt.Println("Failed to connect to the database")
	}

	appAddress := fmt.Sprintf("%s:%d", serverConfig.Application.Host, serverConfig.Application.Port)
	router := routers.SetupRouter(*dbClient, &cfg.Settings.Database)

	// Start server
	server := &http.Server{
		Addr:    appAddress,
		Handler: router,
	}
	fmt.Println(server.Addr)

	app := &application{
		config: *cfg,
		server: server,
	}

	return app, nil
}

func Run() {
	app, err := build()
	if err != nil {
		fmt.Println("Error in build process")
	}
	err = app.server.ListenAndServe()
	if err != nil {
		fmt.Println("Error on Listen and Serve")
	}

}

// DATABASE
func connectDB(dbSettings *config.DatabaseSettings) (*mongo.Client, error) {
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
