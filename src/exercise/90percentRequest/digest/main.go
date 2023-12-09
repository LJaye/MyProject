package main

import (
	"fmt"
	"time"

	"github.com/influxdata/tdigest"
)

type RequestMonitor struct {
	tdigest *tdigest.TDigest
}

func NewRequestMonitor() *RequestMonitor {
	return &RequestMonitor{
		tdigest: tdigest.NewWithCompression(100),
	}
}

func (rm *RequestMonitor) handler(duration time.Duration) {
	rm.tdigest.Add(duration.Seconds(), 1)
}

func (rm *RequestMonitor) calculate90thPercentile() time.Duration {
	percentile := 0.9
	return time.Duration(rm.tdigest.Quantile(percentile) * float64(time.Second))
}

func main() {
	rm := NewRequestMonitor()

	// 模拟处理请求的情况
	for i := 0; i < 1000000; i++ {
		duration := time.Duration(i) * time.Microsecond
		rm.handler(duration)
	}

	// 计算并输出90%分位数
	percentile := rm.calculate90thPercentile()
	fmt.Printf("90th Percentile: %v\n", percentile)
}
