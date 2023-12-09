// 业务逻辑
// Author: 691
// Since: 2023-02-15 16:31
// File: src/go_code/customerManage/service/customerService.go
package service

import "study/src/go_code/customerManage/model"

type CustomerService struct {
	customers []model.Customer
	// 当前切片含多少客户，作为新客户的id
	customerNum int
}

func NewCustomerService() *CustomerService {
	// 初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "liu", "女", 20, "15829927100", "jaye_l@sina.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 客户列表
func (c *CustomerService) ShowList() []model.Customer {
	return c.customers
}

// 添加客户
// 必须*CustomerService，若是值拷贝，则添加的时候原先的客户被丢弃
func (c *CustomerService) AddCustomer(customer model.Customer) bool {
	c.customerNum++
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}

// 通过id查找客户，返回id号对应的切片下标，没找到返回-1
func (c *CustomerService) FindById(id int) int {
	index := -1
	for i, customer := range c.customers {
		if id == customer.Id {
			index = i
		}
	}
	return index
}

// 删除客户
func (c *CustomerService) DelCustomer(id int) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	}
	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}

// 修改客户信息
func (c *CustomerService) UpdateCustomer(id int, name string, gender string, age int, phone string, email string) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	}
	c.customers[index].Name = name
	c.customers[index].Gender = gender
	c.customers[index].Age = age
	c.customers[index].Phone = phone
	c.customers[index].Email = email
	return true
}
