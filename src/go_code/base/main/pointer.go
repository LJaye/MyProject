package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println("i的地址：", &i)

	//ptr是一个指针变量，类型是*int，ptr本身的值是&i
	var ptr *int = &i
	fmt.Printf("ptr = %v\n", ptr)
	fmt.Println("ptr的地址：", &ptr)
	fmt.Println("ptr指向的值：", *ptr)

	//通过p去修改num的值
	var num int = 10
	var p *int = &num
	*p = 100
	fmt.Printf("num = %v\n", num)
}
