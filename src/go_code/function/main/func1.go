package main

import "fmt"

func main() {
	num := 20
	//add(num)
	add1(&num)
	fmt.Println("quantile:", num)
}

func add(num int) {
	num += 10
	fmt.Println("add:", num)
}

func add1(num *int) {
	*num += 10
	fmt.Println("add1", *num)
}
