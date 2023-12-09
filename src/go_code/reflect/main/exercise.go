// 使用反射遍历结构体字段，调用结构体方法并获取结构体标签值
// Author: 691
// Since: 2023-05-14 20:33
// File: src/go_code/reflect/quantile/exercise.go
package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (m Monster) Print() {
	fmt.Println("----start----")
	fmt.Println(m)
	fmt.Println("----end----")
}

func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func TestStruct(m interface{}) {
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	k := v.Kind()
	if k != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	// 获取该结构体有几个字段
	num := v.NumField()
	fmt.Printf("struct has %d fields\n", num) // 4
	// 遍历结构体所有字段
	for i := 0; i < num; i++ {
		// 获取结构体的值
		fmt.Printf("Field %d = %v\n", i, v.Field(i))
		// 获取struct标签，需要通过reflect.Type来获取tag标签
		tagValue := t.Field(i).Tag.Get("json")
		if tagValue != "" {
			fmt.Printf("Field %d tag = %v\n", i, tagValue)
		}
	}

	// 获取该结构体有多少方法
	numOfMethod := v.NumMethod()
	fmt.Printf("struct has %d fields\n", numOfMethod)
	// 获取第2个方法并调用
	// 方法排序是按照函数名字，字母排序（ASCII码），因此Print()是第一个方法
	v.Method(1).Call(nil)
	// 调用第一个方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := v.Method(0).Call(params) // res仍然是切片
	fmt.Println("res = ", res[0].Int())
}

func main() {
	monster := Monster{
		Name:  "lion",
		Age:   400,
		Score: 30.8,
	}
	TestStruct(monster)
}
