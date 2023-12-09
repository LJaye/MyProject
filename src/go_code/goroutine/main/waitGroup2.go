package main

import (
	"fmt"
	"sync"
)

func main() {
	// 一写多读

	wg := new(sync.WaitGroup)
	channel := make(chan int, 10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			channel <- i
		}
		close(channel)
	}() // 后面这个括号是传参用的

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range channel {
				fmt.Println(data)
			}
		}()
	}

	wg.Wait()
}
