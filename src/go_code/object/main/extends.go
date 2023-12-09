// 继承-学生考试系统
// Author: 691
// Since: 2023-01-17 15:18
// File: src/go_code/object/extends.go
package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) showScore() {
	fmt.Println("成绩是：", s.Score)
}

type Pupil struct {
	Student // 嵌入了Student的匿名结构体
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中。。。")
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中。。。")
}

func main() {
	pupil := &Pupil{
		Student{
			Name:  "mary",
			Age:   13,
			Score: 90,
		},
	}
	pupil.testing()
	pupil.showScore()
}
