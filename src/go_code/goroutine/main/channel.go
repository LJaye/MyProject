package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	// 写入数据
	intChan <- 10
	intChan <- 20
	intChan <- 30
	// 写入超过管道容量的数据会造成 deadlock !!
	fmt.Println(len(intChan), cap(intChan))

	num := <-intChan
	fmt.Println(num)
	fmt.Println(len(intChan), cap(intChan))
	// 在没有使用协程的情况下，管道数据为空，再取数据会造成 deadlock !!
}
