package main

import (
	"fmt"
	"strconv"
	"study/src/go_code/base/model"
	"unsafe"
)

//var a = "abc"
//var b = 22
//var c = 33
//以上全局变量声明可等价于
var (
	a = "abc"
	b = 22
	c = 33
)

func main() {
	fmt.Println("hello,world!")
	fmt.Println("a")

	//1.指定变量类型，声明后不赋值
	var i int
	i = 10
	fmt.Println("i=", i)
	//2.根据值自行判断变量类型（类型推导）
	var s = "jackson"
	fmt.Println("name=", s)
	//3.省略var，注意:=左侧的变量不应该是已经声明过的
	name := "jenny"
	fmt.Println("name=", name)
	//等价于
	//var name string
	//name = "jenny"
	//fmt.Println("name=", name)

	//多变量声明
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	var n4, n5, n6 = 1, "str", 10.1
	fmt.Println("n4=", n4, "n5=", n5, "n6=", n6)
	n7, n8, n9 := 1, "str", 10.1
	fmt.Println("n7=", n7, "n8=", n8, "n9=", n9)

	//查看变量数据类型，Printf用于格式化输出
	fmt.Printf("n4的数据类型：%T\n", n4) //int
	fmt.Printf("n6的数据类型：%T\n", n6) //float64
	fmt.Printf("n4占用的字节数：%d\n", unsafe.Sizeof(n4))

	//科学计数法
	num1 := 5.1234e2
	num2 := 5.1234e-2
	fmt.Println("num1=", num1, "num2=", num2)

	//字符
	var c1 byte = 'a'
	var c2 int = 22269
	fmt.Printf("c1=%d\n", c1)
	fmt.Printf("c2=%c\n", c2)

	//字符串
	str1 := "abc\n\"abc"
	fmt.Println(str1)
	str2 := `package quantile

import (
	"fmt"
	"unsafe"
)

func quantile() {
	fmt.Println("hello,world!")
	fmt.Println("a")
}
`
	fmt.Println(str2)

	//基本数据类型转string
	var number1 int = 99
	//var b bool = true
	//var char byte = 'l'
	var str string
	//方式1 fmt.Sprintf
	str = fmt.Sprintf("%d", number1)
	fmt.Printf("字符串类型：%T，str=%v\n", str, str)
	//方式2 strconv包中的函数
	var number2 float64 = 23.456
	str = strconv.FormatFloat(number2, 'f', 2, 64)
	fmt.Printf("字符串类型：%T，str=%q\n", str, str)
	var number3 int = 34
	str = strconv.Itoa(number3)
	fmt.Printf("字符串类型：%T，str=%q\n", str, str)

	//string转基本数据类型（使用strconv包的函数
	//若传递参数有误未转成功则按照默认值来处理
	b, _ := strconv.ParseBool("true")
	fmt.Printf("b type：%T，b=%v\n", b, b)
	n10, _ := strconv.ParseInt("123490", 10, 0)
	fmt.Printf("n10 type：%T，n10=%v\n", n10, n10)

	fmt.Println(model.HeroName)
}
