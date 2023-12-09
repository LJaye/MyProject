// Student结构体
// Author: 691
// Since: 2023-01-11 11:36
// File: src/go_code/struct/factory/model/student.go
package model

type student struct {
	Name  string
	score float64 // score首字母小写，则不可导出
}

func NewStudent(name string, score float64) *student {
	// 构造函数，属性的限制if条件写在此处
	// 返回的数据结构创建在堆上，是共享的数据空间，都可以使用
	return &student{
		Name:  name,
		score: score,
	}
}

func (s *student) GetScore() float64 {
	return s.score
}
