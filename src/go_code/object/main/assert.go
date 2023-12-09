// 类型断言
// Author: 691
// Since: 2023-02-04 20:33
// File: src/go_code/object/quantile/assert.go
package main

import "fmt"

type Stu struct {
}

// 编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for i, item := range items {
		switch item.(type) {
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", i, item)
		case int:
			fmt.Printf("第%v个参数是int类型，值是%v\n", i, item)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", i, item)
		case Stu:
			fmt.Printf("第%v个参数是Stu类型，值是%v\n", i, item)
		case *Stu:
			fmt.Printf("第%v个参数是*Stu类型，值是%v\n", i, item)
		default:
			fmt.Printf("第%v个参数类型不确定,%T，值是%v\n", i, item, item)
		}
	}
}

func main() {
	n1 := 23
	n2 := "liu"
	n3 := 2.3
	n4 := true
	n5 := Stu{}
	n6 := &Stu{}
	TypeJudge(n1, n2, n3, n4, n5, n6)
}
