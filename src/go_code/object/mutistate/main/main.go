// 多态
// Author: 691
// Since: 2023-02-04 18:13
// File: src/go_code/object/mutistate/quantile/hello.go
package main

import "fmt"

type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

// 让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println(p.Name + "手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println(p.Name + "手机停止工作")
}

// Phone独有的方法
func (p Phone) Call() {
	fmt.Println(p.Name + "手机通话")
}

type Camera struct {
	Name string
}

// 让Camera实现Usb接口的方法
func (c Camera) Start() {
	fmt.Println(c.Name + "相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println(c.Name + "相机停止工作")
}

type Computer struct {
}

// 编写一个Working方法，接受Usb接口类型变量
// 只要是实现了Usb接口，就是实现了Usb接口里的所有方法
// usb Usb 体现多态
func (c Computer) Working(usb Usb) {
	// 通过接口变量调用
	usb.Start()
	// 如果是Phone变量，除了调用usb接口声明方法外，还需要调用Phone特有方法=》类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	// 定义一个Usb接口数组，可以存放Phone和Camera的接口体变量
	// 体现多态数组
	var usb [3]Usb
	usb[0] = Phone{"小米"}
	usb[1] = Phone{"vivo"}
	usb[2] = Camera{"索尼"}
	fmt.Println(usb)

	// Phone有一个特有的方法Call()，遍历usb数组
	// 如果是Phone变量，除了调用usb接口声明方法外，还需要调用Phone特有方法=》类型断言
	var computer Computer
	for _, v := range usb {
		computer.Working(v)
	}
}
