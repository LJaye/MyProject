// 客户信息结构体
// Author: 691
// Since: 2023-02-15 16:47
// File: src/go_code/customerManage/model/customer.go
package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂模式实例，返回一个Customer实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 工厂模式实例，返回一个Customer实例，不带id
func NewCustomer2(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 显示该用户信息
func (customer Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", customer.Id, customer.Name, customer.Gender, customer.Age, customer.Phone, customer.Email)
}
