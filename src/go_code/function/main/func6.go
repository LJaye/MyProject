package main

import "fmt"

func main() {
	a, b := cal(50, 10)
	fmt.Printf("a=%v, b=%v", a, b)
}

func cal(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
