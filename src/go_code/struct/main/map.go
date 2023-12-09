package main

import (
	"fmt"
	"sort"
)

func main() {
	//1
	var a map[int]string
	a = make(map[int]string, 10)
	a[1] = "jacky"
	a[4] = "tom"
	a[8] = "rose"
	a[3] = "mia"
	//2
	b := make(map[int]string, 10)
	b[1] = "jackson"
	//3
	c := map[string]string{
		"中国": "北京",
		"美国": "华盛顿",
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//查找
	val, ok := c["中国"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("未找到")
	}

	//遍历
	for k, v := range c {
		fmt.Printf("%v:%v\n", k, v)
	}

	//切片
	var d []map[string]string
	d = make([]map[string]string, 1)
	if d[0] == nil {
		d[0] = make(map[string]string, 2)
		d[0]["name"] = "liu"
		d[0]["age"] = "20"
	}
	//这里需要使用切片的append，使得可以动态增长
	item := map[string]string{
		"name": "jack",
		"age":  "11",
	}
	d = append(d, item)
	fmt.Println(d)

	//排序
	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _, v := range keys {
		fmt.Println(a[v])
	}
}
