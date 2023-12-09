// 统计1-8000的数字中，哪些是素数
// Author: 691
// Since: 2023-05-09 16:02
// File: src/go_code/goroutine/quantile/channelPrime.go
package main

import "fmt"

func main() {
	intChan := make(chan int, 1000)   // 存8000个数字
	primeChan := make(chan int, 2000) // 质数
	exitChan := make(chan bool, 4)    // 退出，开4个协程

	// 开启一个协程，向intChan中放入1-8000
	go putNum(intChan)
	// 开启四个协程，从intChan取出数据，并判断是否为素数，如果是，放到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 循环获取exitChan，直到4个协程都执行完
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	// 输出素数
	for prime := range primeChan {
		fmt.Println(prime)
	}

}

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	// 关闭intChan
	close(intChan)
}

func primeNum(intChan, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag bool
		// 判断是否是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = true // 不是素数
				break
			}
		}

		if !flag {
			primeChan <- num
		}
	}
	exitChan <- true
	fmt.Println("协程退出")
}
