package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 0
	for {
		//生成随机数
		//time.Now().UnixNano()返回一个时间戳（纳秒）
		rand.Seed(time.Now().UnixNano()) //需要设置种子
		n := rand.Intn(100)              //[0,100)
		fmt.Printf("n=%d\n", n)
		count++
		if n == 50 {
			break
		}
	}
	fmt.Printf("生成50共用了%d次", count)
}
