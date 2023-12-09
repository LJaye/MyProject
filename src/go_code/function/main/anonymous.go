package main

import "fmt"

var (
	//func1是全局匿名函数
	func1 = func(n1, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//1.在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	res := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(res)

	//2.将匿名函数赋给一个变量(函数类型)，再通过变量调用匿名函数
	a := func(n1, n2 int) int {
		return n1 - n2
	}
	res1 := a(10, 20)
	fmt.Println(res1)

	//3.如果将匿名函数赋给一个全局变量，那么该匿名函数就称为一个全局匿名函数
	res2 := func1(4, 20)
	fmt.Println(res2)
}
