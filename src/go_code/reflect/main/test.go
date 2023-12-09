package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	// 通过反射获取传入变量的type、kind、值
	// 先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println(rType) // int 不是真正的int类型，而是反射的类型
	fmt.Printf("%T\n", rType)
	// 获取reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Println(rValue)        // 100 无法直接进行加减运算
	fmt.Printf("%T\n", rValue) // reflect.Value
	fmt.Println(rValue.Int() + 2)

	// rValue转为interface{}
	iValue := rValue.Interface()
	// 讲interface{}通过断言转为需要的类型
	num := iValue.(int)
	fmt.Println(num + 2)

	// 获取变量的Kind
	kind1 := rType.Kind()
	kind2 := rValue.Kind()
	fmt.Println(kind1)
	fmt.Println(kind2)
}

func main() {
	var num = 100
	reflectTest01(num)
}
