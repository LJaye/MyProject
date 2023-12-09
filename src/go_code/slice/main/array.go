package main

import "fmt"

func main() {
	//四种初始化数组的方式
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{5, 6, 7}
	var arr3 = [...]int{8, 9, 10}
	var arr4 = [...]int{1: 23, 0: 11, 2: 55}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	arr := [3]int{11, 22, 33}
	visit(&arr)
	fmt.Println(arr)
}

func visit(arr *[3]int) {
	(*arr)[0] = 88
}
