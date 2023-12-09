package main

import "fmt"

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}

func sum(n1, n2 int) int {
	//当执行到defer时，会将defer语句压入到独立到defer栈中（与sum栈和main栈不同），暂时不执行
	//当函数执行完毕后，再从defer栈，按照先入后出当方式出栈
	defer fmt.Println("n1=", n1)
	defer fmt.Println("n2=", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("res=", res)
	return res
}
