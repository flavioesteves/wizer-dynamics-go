package app

import (
	"fmt"
	"wizer-dyanamics-go/config"
)

func Run() {
	fmt.Println("Start the application")
	config, err := config.GetConfiguration()
	if err != nil {
		println("Error %w", err)
	}

	fmt.Println("Settings: ", config)
}
