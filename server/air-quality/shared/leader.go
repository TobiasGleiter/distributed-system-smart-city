package shared

import (
	"server/air-quality/models"
)

var (
	Leader int
	NodeID int
	NodeIP string
	Nodes []models.Node
)

func SetLeader(newLeader int) {
    Leader = newLeader
}

func GetLeader() int {
	return Leader
}

func IsLeader() bool {
	return NodeID == Leader
}

func SetNodes(newNodes []models.Node) {
	Nodes = newNodes
}

func GetNodes() []models.Node {
	return Nodes
}


