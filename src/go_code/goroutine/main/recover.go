package main

import (
	"fmt"
)

func main() {
	go sayHello()
	go handle()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello")
	}
}

func handle() {
	defer func() {
		// 捕获 handle() 抛出的panic
		if err := recover(); err != nil {
			fmt.Println("发生错误", err)
		}
	}()
	var errMap map[int]string
	errMap[0] = "golang"
}
