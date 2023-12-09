// 并发操作map
// Author: 691
// Since: 2023-05-09 21:00
// File: src/go_code/goroutine/quantile/multiMap2.go
package main

import (
	"fmt"
	"sync"
)

// sync.Map详解 https://blog.csdn.net/a348752377/article/details/104972194
func main() {
	a := sync.Map{}
	for i := 0; i < 20; i++ {
		a.Store(i, i)
	}

	val, _ := a.Load(1)
	if valInt, ok := val.(int); ok { // 断言
		valInt++
		fmt.Println(valInt)
	}
}
