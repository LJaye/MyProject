package main

import "fmt"

// 线性表的链式存储

// 链表的节点
type Node struct {
	Data int
	Next *Node
}

// 链表
type LinkedList struct {
	Head *Node
}

// 初始化
func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

// 表长
func (l *LinkedList) Length() int {
	length := 0
	current := l.Head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// 按序号查找
func (l *LinkedList) FindIndex(index int) *Node {
	i := 0
	current := l.Head
	for current != nil {
		if index == i {
			return current
		}
		current = current.Next
		i++
	}
	return nil
}

// 按值查找
func (l *LinkedList) Find(x int) *Node {
	current := l.Head
	for current != nil {
		if current.Data == x {
			return current
		}
		current = current.Next
	}
	return nil
}

// 插入
func (l *LinkedList) Insert(x, i int) {
	p := NewNode(x)
	if i > l.Length() {
		return
	}
	// 插入第一个位置特殊处理
	if i == 0 {
		l.Head = p
	} else {
		s := l.FindIndex(i - 1)
		p.Next = s.Next
		s.Next = p
	}
}

// 删除
func (l *LinkedList) Delete(i int) {
	if i == 0 {
		l.Head = l.Head.Next
	}
	current := l.Head
	index := 0
	for current != nil {
		if index == i-1 {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
		index++
	}
}

// 打印
func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	list := NewLinkedList()
	list.Insert(3, 0)
	list.Insert(8, 1)
	list.Insert(2, 2)
	list.Insert(6, 3)
	list.Insert(7, 4)
	list.Print()
	fmt.Println(list.Find(3))
	fmt.Println(list.FindIndex(0))
	fmt.Println(list.Length())
	list.Delete(4)
	list.Print()
}
