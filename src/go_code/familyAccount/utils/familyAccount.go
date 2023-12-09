// 家庭收支软件
// Author: 691
// Since: 2023-02-15 13:04
// File: src/go_code/familyAccount/utils/familyAccount.go
package utils

import "fmt"

type FamilyAccount struct {
	key     string  // 用户输入选项
	loop    bool    // 是否退出for循环
	flag    bool    // 是否有收支行为
	balance float32 // 余额
	money   float32 // 收支金额
	note    string  // 说明
	details string  // 收支详情
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		flag:    false,
		balance: 10000,
		money:   0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说明",
	}
}

// 收支明细
func (account *FamilyAccount) showDetails() {
	fmt.Println("----------当前收支明细记录----------")
	if account.flag {
		fmt.Println(account.details)
	} else {
		fmt.Println("当前没有收支明细，来一笔吧～")
	}
}

// 登记收入
func (account *FamilyAccount) income() {
	fmt.Print("本次收入金额：")
	fmt.Scanln(&account.money)
	account.balance += account.money //修改账户余额
	fmt.Print("本次收入说明：")
	fmt.Scanln(&account.note)
	// 拼接details变量
	account.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", account.balance, account.money, account.note)
	account.flag = true
}

// 登记支出
func (account *FamilyAccount) pay() {
	fmt.Print("本次支出金额：")
	fmt.Scanln(&account.money)
	if account.balance < account.money {
		fmt.Println("余额不足，请重试")
		return
	} else {
		account.balance -= account.money //修改账户余额
	}
	fmt.Print("本次支出说明：")
	fmt.Scanln(&account.note)
	// 拼接details变量
	account.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", account.balance, account.money, account.note)
	account.flag = true
}

// 退出
func (account *FamilyAccount) exit() {
	fmt.Println("确定要退出吗？y/n")
	choice := ""
	for {
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("error!!!")
			return
		}
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("输入有误！！！")
		return
	}
	if choice == "y" {
		account.loop = false
	}
}

// 主菜单
// 退出
func (account *FamilyAccount) MainMenu() {
	// 显示主菜单
	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退   出")
		fmt.Println("            ")
		fmt.Print("请选择（1-4）：")

		_, err := fmt.Scanln(&account.key)
		if err != nil {
			fmt.Println("error!!!")
			return
		}

		switch account.key {
		case "1":
			account.showDetails()
		case "2":
			account.income()
		case "3":
			account.pay()
		case "4":
			account.exit()

		default:
			fmt.Println("请输入正确的选项！！！")
		}

		if !account.loop {
			fmt.Println("已退出家庭收支软件")
			break
		}
	}
}
