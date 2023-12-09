package main

import "fmt"

type myFuncType func(int, int) int

func main() {
	fmt.Println(myFunc(getSum1, 50, 20))
}

func getSum1(n1 int, n2 int) int {
	return n1 + n2
}

func myFunc(funcVar myFuncType, a int, b int) int {
	return funcVar(a, b)
}
