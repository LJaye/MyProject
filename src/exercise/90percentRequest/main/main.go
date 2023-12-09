package main

import (
	"fmt"
	"github.com/influxdata/tdigest"
	"math/rand"
	"sync"
	"time"
)

// 请求信息
type RequestInfo struct {
	Time     time.Time     // 收到该请求的时刻
	Duration time.Duration // 该请求耗时
}

// 请求监控模块
type RequestMonitor struct {
	mu            sync.Mutex
	requests      []*RequestInfo // 环形缓冲区存储请求信息
	nextIndex     int            // 环形缓冲区写入位置
	maxBufferSize int            // 环形缓冲区大小
	TDigest       *tdigest.TDigest
}

func NewRequestMonitor(maxBufferSize int) *RequestMonitor {
	return &RequestMonitor{
		requests:      make([]*RequestInfo, maxBufferSize),
		nextIndex:     0,
		maxBufferSize: maxBufferSize,
		TDigest:       tdigest.New(),
	}
}

// 处理请求
func (rm *RequestMonitor) handler(r *RequestInfo) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	// 记录请求信息到环形缓冲区
	rm.requests[rm.nextIndex] = r
	rm.nextIndex = (rm.nextIndex + 1) % rm.maxBufferSize
}

// 计算近5分钟内90%的请求响应时间
func (rm *RequestMonitor) calculate90thResponse() time.Duration {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	// 计算近5分钟内的请求记录
	var recentRequests []*RequestInfo
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)
	for _, req := range rm.requests {
		if req == nil {
			continue
		}
		if req.Time.After(fiveMinutesAgo) {
			recentRequests = append(recentRequests, req)
			rm.TDigest.Add(req.Duration.Seconds(), 1)
		}
	}

	// 近似算出90%分位数
	return time.Duration(rm.TDigest.Quantile(0.9) * float64(time.Second))

	/*
		// 排序请求记录
		sort.Slice(recentRequests, func(i, j int) bool {
			return recentRequests[i].Duration < recentRequests[j].Duration
		})

		// 计算90%分位数
		index := int(0.9 * float64(len(recentRequests)))
		if recentRequests == nil {
			return 0
		}

		return recentRequests[index].Duration
	*/
}
func main() {
	// 创建请求监控实例
	monitor := NewRequestMonitor(100)
	for true {
		// 模拟请求
		go func() {
			rand.Seed(time.Now().UnixNano())
			duration := rand.Intn(50) + 50 // 模拟请求随机用时50~100ms
			monitor.handler(&RequestInfo{
				Time:     time.Now(),
				Duration: time.Duration(duration) * time.Millisecond,
			})
		}()
		// 监控结果
		go func() {
			t := monitor.calculate90thResponse()
			fmt.Printf(time.Now().String()+"  近5分钟内90%%的请求响应时间小于：%s\n", t)
		}()
	}
}
