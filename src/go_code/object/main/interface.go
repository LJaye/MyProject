// 面向接口编程
// Author: 691
// Since: 2023-01-18 10:37
// File: src/go_code/object/quantile/interface.go
package main

import "fmt"

type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

// 让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
}

// 让Camera实现Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

// 编写一个Working方法，接受Usb接口类型变量
// 只要是实现了Usb接口，就是实现了Usb接口里的所有方法
// usb Usb 体现多态
func (c Computer) Working(usb Usb) {
	// 通过接口变量调用
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	//phone := Phone{}
	camera := Camera{}

	computer.Working(camera)
}
