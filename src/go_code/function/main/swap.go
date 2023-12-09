package main

import "fmt"

func main() {
	n1 := 10
	n2 := 20
	swap(&n1, &n2)
	fmt.Printf("n1=%v, n2=%v", n1, n2)
}

func swap(n1, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}
