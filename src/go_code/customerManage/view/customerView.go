// 客户关系软件界面
// Author: 691
// Since: 2023-02-15 16:05
// File: src/go_code/customerManage/view/customerView.go
package view

import (
	"fmt"
	"study/src/go_code/customerManage/model"
	"study/src/go_code/customerManage/service"
)

type CustomerView struct {
	Key             string // 用户输入选项
	Loop            bool   // 是否循环显示主菜单
	CustomerService *service.CustomerService
}

// 显示主菜单
func (cv *CustomerView) Mainmenu() {
	for {
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退   出")
		fmt.Println("            ")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&cv.Key)
		switch cv.Key {
		case "1":
			cv.addCustomer()
		case "2":
			cv.updateCustomer()
		case "3":
			cv.delCustomer()
		case "4":
			cv.showList()
		case "5":
			cv.exit()

		default:
			fmt.Println("请输入正确的选项！！！")
		}

		if !cv.Loop {
			fmt.Println("已退出客户关系管理系统")
			break
		}
	}
}

// 显示所有客户信息
func (cv *CustomerView) showList() {
	fmt.Println("----------客户列表----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, customer := range cv.CustomerService.ShowList() {
		fmt.Println(customer.GetInfo())
	}
}

// 添加客户
func (cv *CustomerView) addCustomer() {
	var name, gender, phone, email string
	var age int
	fmt.Println("----------添加客户----------")
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("电话：")
	fmt.Scanln(&phone)
	fmt.Print("邮件：")
	fmt.Scanln(&email)

	// id号没有让用户直接输入，系统自动分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if cv.CustomerService.AddCustomer(customer) {
		fmt.Println("----------添加完成----------")
	} else {
		fmt.Println("----------添加失败----------")
	}
}

// 删除客户
func (cv *CustomerView) delCustomer() {
	fmt.Println("----------删除客户----------")
	fmt.Print("请选择待删除的客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	choice := ""
	for {
		fmt.Println("确定要删除吗？y/n")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("error!!!")
			return
		}
		if choice == "y" || choice == "n" || choice == "Y" || choice == "N" {
			break
		}
		fmt.Println("输入有误！！！")
	}

	if choice == "y" || choice == "Y" {
		if cv.CustomerService.DelCustomer(id) {
			fmt.Println("----------删除完成----------")
		} else {
			fmt.Println("----------删除失败----------")
		}
	}

}

// 退出
func (cv *CustomerView) exit() {
	choice := ""
	for {
		fmt.Println("确定要退出系统吗？y/n")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("error!!!")
			return
		}
		if choice == "y" || choice == "n" || choice == "Y" || choice == "N" {
			break
		}
		fmt.Println("输入有误！！！")
	}

	if choice == "y" || choice == "Y" {
		cv.Loop = false
	}
}

// 修改客户信息
func (cv *CustomerView) updateCustomer() {
	fmt.Println("----------修改客户----------")
	fmt.Print("请选择待删除的客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	var name, gender, phone, email string
	var age int
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("电话：")
	fmt.Scanln(&phone)
	fmt.Print("邮件：")
	fmt.Scanln(&email)

	if cv.CustomerService.UpdateCustomer(id, name, gender, age, phone, email) {
		fmt.Println("----------修改完成----------")
	} else {
		fmt.Println("----------修改失败----------")
	}

}
