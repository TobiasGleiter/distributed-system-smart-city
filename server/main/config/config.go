package config

import (
	"os"
	"encoding/json"

	"server/main/models"
)

type Config struct {
    ID      string `json:"ID"`
    Coordinator bool   `json:"coordinator"`
	Port string `json:"port"`
    MongoURI string `json:"mongo_uri"`
    Nodes []models.Node `json:"nodes"`
}

func LoadConfig(filename string) (Config, error) {
    var config Config
    configFile, err := os.Open(filename)
    if err != nil {
        return config, err
    }
    defer configFile.Close()

    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    return config, err
}