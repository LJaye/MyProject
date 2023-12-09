package main

import "fmt"

func writeData(intChan chan int) {
	// 写数据
	for i := 0; i < 20; i++ {
		intChan <- i
		fmt.Println("write", i)
	}
	// 关闭管道
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	// 遍历数据
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read", v)
	}
	// 任务完成，发送退出信号
	exitChan <- true
	close(exitChan)
}

func main() {
	// 创建两个管道
	intChan := make(chan int, 20)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)

	for v := range exitChan {
		if v == true {
			break
		}
	}
}
