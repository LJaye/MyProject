package main

import "fmt"

// 线性表的顺序存储
type LinearList struct {
	Data []int
}

// 初始化
func NewLinearList() *LinearList {
	return &LinearList{}
}

// 获取线性表的长度
func (l *LinearList) Length() int {
	return len(l.Data)
}

// 查找 成功的平均比较次数为(n+1)/2 运气好第一个找到 运气不好最后一个找到 平均一下除以2
func (l *LinearList) Find(x int) int {
	index := 0
	for index < l.Length() {
		if x == l.Data[index] {
			return index
		}
		index++
	}
	return -1
}

// 插入
func (l *LinearList) Insert(x, i int) {
	if i < 0 || i > l.Length() {
		return
	}
	// 应该插入i-1的位置
	// l.Data[:i]表示0～i-1的元素 l.Data[i:]表示从i～末尾
	l.Data = append(l.Data[:i], append([]int{x}, l.Data[i:]...)...)
}

// 删除
func (l *LinearList) Delete(i int) {
	if i < 0 || i >= l.Length() {
		return
	}
	l.Data = append(l.Data[:i], l.Data[i+1:]...)
}

func main() {
	list := &LinearList{Data: []int{3, 5, 2, 7}}
	fmt.Println(list.Find(5))
	list.Insert(4, 0) // 4,3,5,2,7
	fmt.Println(list)
	list.Delete(4) // 4,3,5,2
	fmt.Println(list)
}
