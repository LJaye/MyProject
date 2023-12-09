package main

import (
	"fmt"
	"sync"
	"time"
)

/**
计算1-200各个数的阶乘，并且把各个数的阶乘放到map中
*/

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	lock sync.Mutex
)

func factorial(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	// 开启多个协程
	for i := 0; i < 20; i++ {
		go factorial(i)
	}
	time.Sleep(time.Second * 10)
	// 如果以上协程在10秒内没有执行完，还可能出现资源争夺，则读的时候还需要加锁
	// lock.Lock()
	fmt.Println(myMap)
	// lock.Unlock()
}
