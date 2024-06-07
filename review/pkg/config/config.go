package config

import (
	"encoding/json"
	"github.com/brumhard/alligotor"
	"log"
	"os"
)

// For running locally
//var filePath = "./config/config.json"

// Uncomment the below variable and comment the upper variable when creating a docker image
// IMPORTANT: That file has to be mounted as a volume when running the docker container
//var filePath = "/go/src/coupon-service/config/config.json"
var filePath = "../config/config.json"

type ApiConfig struct {
	Host string
	Port int
}

type Config struct {
	Api *ApiConfig
}

func loadConfig(filePath string) (Config, error) {
	var config Config
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config, nil
}

func NewConfig() *Config {
	config, err := loadConfig(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if err := alligotor.Get(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}
