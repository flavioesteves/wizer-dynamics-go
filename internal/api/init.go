package api

import (
	"fmt"
	"net/http"

	"github.com/flavioesteves/wizer-dynamics-go/configs"
	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/flavioesteves/wizer-dynamics-go/internal/routers"
)

type application struct {
	server *http.Server
}

func build() (*application, error) {

	serverConfig, err := config.GetConfiguration()
	if err != nil {
		fmt.Println("error:", err)
	}

	dbClient, e := db.ConnectDB(&serverConfig.Database)
	if e != nil {
		fmt.Println("Failed to connect to the database")
	}

	appAddress := fmt.Sprintf("%s:%d", serverConfig.Application.Host, serverConfig.Application.Port)
	router := routers.SetupRouter(dbClient.Database(serverConfig.Database.DatabaseName))

	// Start server
	server := &http.Server{
		Addr:    appAddress,
		Handler: router,
	}

	return &application{server: server}, nil
}

func Run() {
	app, err := build()
	if err != nil {
		fmt.Println("Error in build process")
	}

	fmt.Println("Server start at: ", app.server.Addr)
	err = app.server.ListenAndServe()
	if err != nil {
		fmt.Println("Error on Listen and Serve")
	}
}
