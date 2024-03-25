package cpu

import (
	"runtime"
	"time"
	"fmt"
)

var (
	threshold = 90
)

func SetThreshold(newThreshold int) {
	threshold = newThreshold
}

func Monitor() {
	for {
		cpuUsage := runtime.NumCPU()

		fmt.Printf(fmt.Sprintf("CPU: %d\n", cpuUsage))

        time.Sleep(time.Second * 10)
	}
}