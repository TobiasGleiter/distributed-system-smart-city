package cpu

import (
	//"runtime"
	"time"
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

type Stats struct{}

func (s *Stats) GetCPUUsage() (float64, error) {
   cpuPercentages, err := cpu.Percent(time.Second, false)
   if err != nil {
	   return 0, err
   }

   var totalCpuUsage float64
   for _, usage := range cpuPercentages {
	   totalCpuUsage += usage
   }
   averageCpuUsage := totalCpuUsage / float64(len(cpuPercentages))

   fmt.Println("CPU usage:", averageCpuUsage)

   return averageCpuUsage, nil
}