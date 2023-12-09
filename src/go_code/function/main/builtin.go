package main

import "fmt"

func main() {
	n1 := 100
	fmt.Printf("n1类型为%T，n1值为%v，n1地址为%v\n", n1, n1, &n1)
	n2 := new(int)
	*n2 = 50
	fmt.Printf("n2类型为%T，n2值为%v，n2地址为%v，n2指向的值为%v", n2, n2, &n2, *n2)
}
