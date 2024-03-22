package models

type Node struct {
	ID int    `json:"ID"`
	IP string `json:"ip"`
}

type Config struct {
	ID       int    `json:"ID"`
	Leader   int    `json:"leader"`
	IP       string `json:"ip"`
	MongoURI string `json:"mongo_uri"`
	Nodes    []Node`json:"nodes"`
}
