package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	// 传统方法在遍历管道时，如果不关闭会阻塞导致deadlock
	// 实际开发中，可能不好确认何时关闭管道
	for {
		select {
		case v := <-intChan: // 如果intChan未关闭，不会一直阻塞，会向下执行
			fmt.Println(v)
		case v := <-stringChan:
			fmt.Println(v)
		default:
			fmt.Println("都取不到，可以加入自己的逻辑")
			return
		}
	}
}
