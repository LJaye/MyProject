package main

import (
	"fmt"
	"go/types"
	"math/rand"
	"time"
)

func main() {
	//switch判断某个interface变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case types.Nil:
		fmt.Printf("x的类型：%T\n", i)
	case int:
		fmt.Println("x的类型：int")
	case float64:
		fmt.Println("x的类型：float64")
	default:
		fmt.Println("未知型")
	}

	//使用for-range遍历字符串
	var str string = "hello北京"
	//输出中文需要将str转成rune类型切片（一个汉字占三个字节
	str1 := []rune(str)
	for index, val := range str1 {
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
	//for-range可以直接输出中文
	str = "hello上海"
	for index, val := range str {
		fmt.Printf("index=%d,val=%c\n", index, val)
	}

	//break默认会跳出最近的for循环，也可以指定标签跳出其对的for循环
label:
	for i := 1; i <= 4; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				break label
			}
			fmt.Printf("j=%d\n", j)
		}
	}

	//生成随机数
	//time.Now().UnixNano()返回一个时间戳（纳秒）
	rand.Seed(time.Now().UnixNano()) //需要设置种子
	n := rand.Intn(100)          //[0,100)
	fmt.Println(n)

	//接收用户输入数据
	var name string
	var age byte
	var salary float32
	var isPass bool
	//fmt.Println("请输入姓名：")
	//fmt.Scanln(&name) //传地址，函数内部修改会影响外部的变量
	//fmt.Println("请输入年龄：")
	//fmt.Scanln(&age)
	//fmt.Println("请输入薪水：")
	//fmt.Scanln(&salary)
	//fmt.Println("请输入是否通过：")
	//fmt.Scanln(&isPass)
	//fmt.Printf("姓名：%v\n年龄：%v\n薪水：%v\n是否通过：%v\n", name, age, salary, isPass)
	//fmt.Scanf可以按指定的格式输入
	fmt.Println("请输入姓名，年龄，薪水，是否通过（使用空格隔开）")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	fmt.Printf("姓名：%v\n年龄：%v\n薪水：%v\n是否通过：%v\n", name, age, salary, isPass)
}
