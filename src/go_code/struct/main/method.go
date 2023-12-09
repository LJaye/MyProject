package main

import "fmt"

type Person struct {
	Name string
}

//给Person类型绑定一个方法
func (p Person) getName() {
	p.Name = "jenny"
	fmt.Println(p.Name)
}

func (p Person) test1() {
	p.Name = "jenny"
	fmt.Println("test1()=", p.Name) // jenny
}

func (p *Person) test2() {
	p.Name = "mary"
	fmt.Println("test2()=", p.Name) // mary
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	return fmt.Sprintf("Name=[%v] Age=[%v]", s.Name, s.Age)
}

func main() {
	//var p Person
	//p.Name = "joye"
	//p.getName()
	//fmt.Println(p.Name)
	//
	//student := &Student{
	//	Name: "liu",
	//	Age:  23,
	//}
	//fmt.Println(student)

	person := Person{"tom"}
	person.test1()
	fmt.Println("quantile()=", person.Name) // tom
	(&person).test1()                       // 从形式上传入是地址，但本质上还是值拷贝
	fmt.Println("quantile()=", person.Name) // tom
	(&person).test2()
	fmt.Println("quantile()=", person.Name) // mary
	person.test2()                          // 编译器优化处理
	fmt.Println("quantile()=", person.Name) // mary
}
