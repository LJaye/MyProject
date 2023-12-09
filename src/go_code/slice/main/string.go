package main

import "fmt"

func main() {
	str := "jack"
	b := []rune(str)
	b[0] = '刘'
	str = string(b)
	fmt.Println(str)
}
