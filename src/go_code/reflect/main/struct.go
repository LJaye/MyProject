// 对结构体进行操作
// Author: 691
// Since: 2023-05-14 14:21
// File: src/go_code/reflect/quantile/struct.go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	// 通过反射获取传入变量的type、kind、值
	// 先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println(rType)
	fmt.Printf("%T\n", rType)
	// 获取reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Println(rValue)

	// rValue转为interface{}
	iValue := rValue.Interface()
	fmt.Printf("%v,%T\n", iValue, iValue)
	// 反射的本质是运行时获取信息，所以并不能通过iValue直接获取结构体的属性，如果iValue.Name则编译时就会报错
	// 讲interface{}通过断言转为需要的类型
	stu, ok := iValue.(Student)
	if ok {
		fmt.Println(stu.Name)
	}
}

func main() {
	stu := Student{
		Name: "joy",
		Age:  24,
	}
	reflectTest02(stu)
}
