package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	now := time.Now() //类型是time.Time
	fmt.Println(now)  //2021-11-13 17:11:06.013748 +0800 CST m=+0.000051293
	//获取年月日时分秒
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))
	//格式化日期和时间
	//1.Printf或Sprintf
	fmt.Printf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	//2.now.Format
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	fmt.Println("start")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("end")
}
