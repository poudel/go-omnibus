package systeminfo

import (
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

type Metric struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

func ServeMetrics(ch chan<- Metric) {
	for {
		now := time.Now()

		v, _ := mem.VirtualMemory()
		r := Metric{Name: "AvailableRam", Value: float64(v.Available) / 1e6, Timestamp: now.Unix()}

		ch <- r

		time.Sleep(3 * time.Second)
	}
}
