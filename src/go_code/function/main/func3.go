package main

import "fmt"

func main() {
	a := getSum
	fmt.Printf("a的类型%T，getSum类型%T\n", a, getSum)
	fmt.Println(a(10, 20))
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}
