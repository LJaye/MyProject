// json反序列化
// Author: 691
// Since: 2023-03-04 17:28
// File: src/go_code/file/quantile/unserialize.go
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string `json:"name"` // tag使用到反射机制
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	str := "{\"name\":\"liu\",\"age\":23,\"address\":\"liaoning\"}"
	var s Student
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	// map不需要make，因为Unmarshal中有封装make操作
	var s1 map[string]interface{}
	json.Unmarshal([]byte(str), &s1)
	fmt.Println(s1)
}
