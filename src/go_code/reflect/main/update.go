// 通过反射修改值
// Author: 691
// Since: 2023-05-14 20:18
// File: src/go_code/reflect/quantile/update.go
package main

import (
	"fmt"
	"reflect"
)

func reflectTest03(b interface{}) {
	rValue := reflect.ValueOf(b)
	fmt.Println(rValue.Kind()) // ptr num传入的是地址，因此是指针类型
	rValue.Elem().SetInt(200)
	// Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
	// 如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。
}

func main() {
	num := 100
	reflectTest03(&num)
	fmt.Println(num)
}
