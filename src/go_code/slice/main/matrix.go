package main

import "fmt"

func main() {
	var arr = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr)
	for i, v := range arr {
		for j, v1 := range v {
			fmt.Printf("arr[%v][%v]=%v  ", i, j, v1)
		}
		fmt.Println()
	}
}
