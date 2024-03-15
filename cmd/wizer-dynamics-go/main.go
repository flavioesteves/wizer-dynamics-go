package main

import (
	"fmt"
	"github.com/flavioesteves/wizer-dynamics-go/configs"
	"github.com/flavioesteves/wizer-dynamics-go/internal/app"
)

func main() {
	configuration, err := config.GetConfiguration()
	if err != nil {
		fmt.Println("error:", err)
	}

	appBuild, err := app.Build(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	appBuild.Server.ListenAndServe()
}
