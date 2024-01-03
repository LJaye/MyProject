package main

import "fmt"

// 双向链表

type DoubleNode struct {
	Data int
	Prev *DoubleNode
	Next *DoubleNode
}

type DoubleLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

func NewDoubleNode(data int) *DoubleNode {
	return &DoubleNode{Data: data, Prev: nil, Next: nil}
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{Head: nil, Tail: nil}
}

func (d *DoubleLinkedList) Length() int {
	p := d.Head
	i := 0
	for p != nil {
		p = p.Next
		i++
	}
	return i
}

func (d *DoubleLinkedList) FindIndex(i int) *DoubleNode {
	if i > d.Length() {
		return nil
	}
	p := d.Head
	for j := 0; j < i; j++ {
		p = p.Next
	}
	return p
}

func (d *DoubleLinkedList) InsertAtTail(x int) {
	node := NewDoubleNode(x)
	if d.Length() == 0 {
		d.Head = node
		d.Tail = node
	} else {
		d.Tail.Next = node
		node.Prev = d.Tail
		d.Tail = d.Tail.Next
	}
}

func (d *DoubleLinkedList) DeleteAtHead() {
	if d.Length() == 0 {
		return
	}
	d.Head = d.Head.Next
}

func (d *DoubleLinkedList) Print() {
	p := d.Head
	for p != nil {
		fmt.Print(p.Data, " ")
		p = p.Next
	}
	fmt.Println()
}

func main() {
	list := NewDoubleLinkedList()
	list.InsertAtTail(1)
	list.InsertAtTail(4)
	list.InsertAtTail(3)
	list.InsertAtTail(8)
	fmt.Println(list.Length())
	list.Print()
	list.DeleteAtHead()
	list.DeleteAtHead()
	list.Print()
}
