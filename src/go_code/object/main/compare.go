// 继承vs接口
// Author: 691
// Since: 2023-02-04 15:21
// File: src/go_code/object/quantile/compare.go
package main

import "fmt"

type Monkey struct {
	Name string
}

func (monkey *Monkey) Climb() {
	fmt.Println(monkey.Name + "爬树")
}

// 继承Monkey
type LittleMonkey struct {
	Monkey
}

type BirdAble interface {
	Fly()
}

type FishAble interface {
	Swim()
}

// 实现BirdAble接口
func (littleMonkey *LittleMonkey) Fly() {
	fmt.Println(littleMonkey.Name + "飞翔")
}

// 实现FishAble接口
func (littleMonkey *LittleMonkey) Swim() {
	fmt.Println(littleMonkey.Name + "游泳")
}

func main() {
	m := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	m.Climb()
	m.Fly()
	m.Swim()
}
