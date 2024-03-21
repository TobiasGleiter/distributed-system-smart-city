package config

import (
	"os"
	"encoding/json"
)

type Node struct {
    Port string `json:"port"`
}

type Config struct {
	Port string `json:"port"`
    MongoURI string `json:"mongo_uri"`
    Nodes []Node `json:"nodes"`
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