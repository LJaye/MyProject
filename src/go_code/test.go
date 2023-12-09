package main

import "fmt"

func main() {
	d := []int{123, 567}
	e := []string{"lalala", "www"}
	for _, i := range d {
		for _, j := range e {
			if (i == 789 || i == 123) && j == "lalala" {
				fmt.Println(i)
				fmt.Println(j)
				fmt.Println(111)
				fmt.Println("---")
			} else {
				fmt.Println(i)
				fmt.Println(j)
				fmt.Println(222)
				fmt.Println("---")
			}
		}

	}

}
