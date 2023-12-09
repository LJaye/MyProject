// quantile
// Author: 691
// Since: 2023-01-11 11:31
// File: src/go_code/struct/factory/quantile/hello.go
package main

import (
	"fmt"
	"study/src/go_code/struct/factory/model"
)

func main() {
	//stu := model.Student{
	//	Name:  "jaye",
	//	Score: 99.9,
	//}
	//fmt.Println(stu)

	stu := model.NewStudent("jaye", 99.9)
	fmt.Println(stu) // 如果不希望展示&，fmt.Println(*stu)
	fmt.Println(stu.GetScore())
}
