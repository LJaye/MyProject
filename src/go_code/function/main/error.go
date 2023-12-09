package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	fmt.Println("test后面的代码")
	test1()
	fmt.Println("test1后面的代码")
}

func test() {
	defer func() {
		//内置函数recover()来捕获异常
		if err := recover(); err != nil { //说明捕获到错误
			fmt.Println("err=", err)
			//这里可以将错误发送给管理员
			fmt.Println("发送邮件到jingyi20@staff.weibo.com")
		}
	}()
	n1 := 10
	n2 := 0
	res := n1 / n2
	fmt.Println("res=", res)
}

//如果文件名传入不正确，返回自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		fmt.Println("配置文件读取中")
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test1() {
	err := readConf("config.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("继续执行")
}
