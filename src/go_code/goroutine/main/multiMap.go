// 并发操作map
// Author: 691
// Since: 2023-05-09 21:49
// File: src/go_code/goroutine/quantile/multiMap.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(map[int]int, 100)
	var lock sync.Mutex
	for i := 0; i < 100; i++ {
		x := i
		//go func() {
		lock.Lock()
		a[x] = x
		lock.Unlock()
		//}()
	}
	fmt.Println(a)
}
