package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Define your configuration struct here
type Settings struct {
	Application ApplicationSettings `yaml:"application"`
	Database    DatabaseSettings    `yaml:"database"`
	Cloud       CloudSettings       `yaml:"cloud"`
	JWT         JWTSettings         `yaml:"jwt"`
	Environment string
}

type ApplicationSettings struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type DatabaseSettings struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Port         int    `yaml:"port"`
	Host         string `yaml:"host"`
	Model        string `yaml:"model"`
	DatabaseName string `yaml:"database_name"`
}

type JWTSettings struct {
	Secret    string `yaml:"secret"`
	ExpiredIn string `yaml:"expired_in"`
	MaxAge    int    `yaml:"max_age"`
}

type CloudSettings struct {
	MongodbURI string `yaml:"mongodb_uri"`
}

func GetConfiguration() (Settings, error) {
	var settings Settings

	workingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// Build the config files path
	basePath := fmt.Sprintf("%s/config-files/base.yaml", workingDirectory)
	localPath := fmt.Sprintf("%s/config-files/local.yaml", workingDirectory)
	productionPath := fmt.Sprintf("%s/config-files/production.yaml", workingDirectory)

	env := os.Getenv("APP_ENVIRONMENT")
	if env == "" {
		env = "local"
	}

	baseFile, err := os.Open(basePath)
	if err != nil {
		fmt.Println("Error opening base.yaml", err)
	}
	defer baseFile.Close()

	decoder := yaml.NewDecoder(baseFile)
	err = decoder.Decode(&settings)
	if err != nil {
		fmt.Println("Error decoding base.yaml", err)
	}

	var envFile *os.File

	switch env {
	case "local":
		envFile, err = os.Open(localPath)
		settings.Environment = "local"
	case "production":
		envFile, err = os.Open(productionPath)
		settings.Environment = "production"
	default:
		fmt.Println("Warning Unknown environment", env)
	}

	if err != nil && err != os.ErrNotExist { // Ignore "file not found" error
		fmt.Println("Error opening", env, ".yaml:", err)
	}

	if envFile != nil {
		defer envFile.Close()
		decoder := yaml.NewDecoder(envFile)
		err = decoder.Decode(&settings)
		if err != nil {
			fmt.Println("Error decoding", env, ".yaml:", err)
		}
	}

	return settings, nil
}
