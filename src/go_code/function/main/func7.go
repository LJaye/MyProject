package main

import "fmt"

func main() {
	fmt.Println(get0ToManySum(2,3))
	fmt.Println(get0ToManySum(2,3))
}

//求0个到多个参数的和
func get0ToManySum(args ...int) int {
	var sum int
	//遍历args
	for _, v := range args {
		sum += v
	}
	return sum
}

//求1个到多个参数的和
func get1ToManySum(n1 int, args ...int) int {
	sum := n1
	//遍历args
	for _, v := range args {
		sum += v
	}
	return sum
}
