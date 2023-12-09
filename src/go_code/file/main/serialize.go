// json序列化
// Author: 691
// Since: 2023-03-04 16:00
// File: src/go_code/file/quantile/serialize.go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"` // tag使用到反射机制
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	p := Person{
		Name:    "liu",
		Age:     23,
		Address: "liaoning",
	}
	// 结构体序列化
	data, _ := json.Marshal(&p)
	fmt.Println(string(data))

	// map序列化
	a := make(map[string]int, 10)
	a["age"] = 23
	a["salary"] = 18000
	data, _ = json.Marshal(a)
	fmt.Println(string(data))

	// 切片序列化
	var b []map[string]int
	m1 := make(map[string]int, 10)
	m1["age"] = 23
	m1["salary"] = 18000
	m2 := make(map[string]int, 10)
	m2["score"] = 100
	m2["date"] = 20230304
	b = append(b, m1)
	b = append(b, m2)
	data, _ = json.Marshal(b)
	fmt.Println(string(data))
}
