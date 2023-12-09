package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

func main() {
	//1
	var cat1 Cat //声明过后内存空间就已经分配了
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println(cat1)
	//2
	var cat2 = Cat{
		"喵",
		3,
		"橘黄",
	}
	fmt.Println(cat2)
	//3
	var cat3 *Cat = new(Cat)
	cat3.Name = "jk" //(*cat3).Name="jk"，为了使用方便，go的设计者对底层进行优化
	cat3.Age = 1
	cat3.Color = "black"
	fmt.Println(*cat3)
	//4
	var cat4 *Cat = &Cat{
		"rose",
		6,
		"white",
	}
	fmt.Println(*cat4)

	var cat5 *Cat = &cat1
	cat5.Name = "qq"
	fmt.Println(cat1)
	fmt.Println(*cat5)

	cat6 := Cat{"tommy", 5, "yellow"}
	//将cat6变量序列化为一个json字符串
	str, err := json.Marshal(cat6)
	if err != nil {
		fmt.Println("处理错误")
	}
	fmt.Println(string(str))
}
