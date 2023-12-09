// 接口实践 - 接口体切片排序
// Author: 691
// Since: 2023-02-04 15:30
// File: src/go_code/object/quantile/exercise.go
package main

import (
	"fmt"
	"sort"
)

// 1、声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2、声明一个Hero结构体切片类型
type HeroSlice []Hero

// 3、实现Interface接口
/*
type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
}
*/

func (hero HeroSlice) Len() int {
	return len(hero)
}

// 按年龄从小到大
func (hero HeroSlice) Less(i, j int) bool {
	return hero[i].Age < hero[j].Age
}

func (hero HeroSlice) Swap(i, j int) {
	hero[i], hero[j] = hero[j], hero[i]
}

func main() {
	heros := HeroSlice{
		{
			Name: "liu",
			Age:  10,
		},
		{
			Name: "zhang",
			Age:  4,
		},
		{
			Name: "tom",
			Age:  22,
		},
		{
			Name: "jenny",
			Age:  18,
		},
	}
	fmt.Println(heros)
	sort.Sort(heros)
	fmt.Println(heros)
}
