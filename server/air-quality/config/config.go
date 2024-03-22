package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Node struct {
	ID       int    `json:"id"`
	Leader   int    `json:"leader"`
	IP       string `json:"ip"`
	MongoURI string `json:"mongo_uri"`
	Nodes    []struct {
		ID int    `json:"id"`
		IP string `json:"ip"`
	} `json:"nodes"`
}

func LoadConfig(filename string) (*Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Node
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}