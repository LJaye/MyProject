// 设置cpu核数
// Author: 691
// Since: 2023-04-22 21:51
// File: src/go_code/goroutine/quantile/cpu.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	// 可以设置使用cpu核数
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
