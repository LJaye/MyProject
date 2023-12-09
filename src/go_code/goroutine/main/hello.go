// 协程
// Author: 691
// Since: 2023-04-22 20:45
// File: src/go_code/goroutine/quantile/hello.go
package main

import (
	"fmt"
	"time"
)

/**
主线程中（可以理解为进程）中，开启一个goroutine，该协程每一秒输出一个"hello world"
在主线程中每隔一秒输出"hello golang"，输出10次后退出程序
*/

func helloWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
		time.Sleep(time.Second)
	}
}

func main() {
	// 开启一个协程
	go helloWorld()

	for i := 0; i < 10; i++ {
		fmt.Println("hello golang")
		time.Sleep(time.Second)
	}
}
