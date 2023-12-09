package main

import "fmt"

func main() {
	//1.让切片引用一个创建好的数组
	arr := [...]int{1, 22, 33, 66, 99}
	slice := arr[1:3] //引用数组的第2个元素到第3个元素
	fmt.Println(slice)
	fmt.Println("切片的容量", cap(slice))
	slice[1] = 34
	fmt.Println(arr)
	fmt.Println(slice)
	//2.通过make来创建切片
	var slice1 = make([]int, 5, 10) //容量可选
	slice1[2] = 4
	fmt.Println(slice1)
	//3.直接指定具体数组
	var slice2 = []string{"jackson", "rose", "tom"}
	fmt.Println(slice2)
	fmt.Println("切片的长度", len(slice2))
	fmt.Println("切片的容量", cap(slice2))

	var slice3 = []int{100, 200, 300}
	slice3 = append(slice3, 400, 500)
	slice3 = append(slice3, slice1...) //也可以追加切片
	fmt.Println(slice3)

	var a = []int{1, 2, 3}
	var slice4 = make([]int, 10)
	copy(slice4, a)
	fmt.Println(a)
	fmt.Println(slice4)
}
