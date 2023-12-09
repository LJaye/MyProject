package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

// 请求记录结构体
type RequestRecord struct {
	Time     time.Time
	Duration time.Duration
}

// 请求监控模块
type RequestMonitor struct {
	mu            sync.Mutex
	requests      []RequestRecord
	requestsCount int
}

// 处理HTTP请求的处理函数
func (rm *RequestMonitor) handler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	// 处理请求的业务逻辑
	// ...

	duration := time.Since(startTime)

	rm.mu.Lock()
	defer rm.mu.Unlock()

	// 记录请求信息
	rm.requests = append(rm.requests, RequestRecord{
		Time:     time.Now(),
		Duration: duration,
	})

	rm.requestsCount++
}

// 计算近5分钟内90%的请求响应时间
func (rm *RequestMonitor) calculate90thPercentile() time.Duration {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	// 计算近5分钟内的请求记录
	var recentRequests []RequestRecord
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)
	for _, req := range rm.requests {
		if req.Time.After(fiveMinutesAgo) {
			recentRequests = append(recentRequests, req)
		}
	}

	// 排序请求记录
	sort.Slice(recentRequests, func(i, j int) bool {
		return recentRequests[i].Duration < recentRequests[j].Duration
	})

	// 计算90%分位数
	index := int(0.9 * float64(len(recentRequests)))
	if index >= len(recentRequests) {
		index = len(recentRequests) - 1
	}

	return recentRequests[index].Duration
}

func main() {
	// 创建请求监控实例
	requestMonitor := &RequestMonitor{}

	// 启动HTTP服务器
	http.HandleFunc("/", requestMonitor.handler)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("HTTP server error:", err)
		}
	}()

	// 定期输出近5分钟内90%的请求响应时间
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		duration := requestMonitor.calculate90thPercentile()
		fmt.Printf("近5分钟内90%%的请求响应时间小于：%s\n", duration)
	}
}
