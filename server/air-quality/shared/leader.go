// In shared/leader.go

package shared

var (
	Leader int
	NodeID int
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


