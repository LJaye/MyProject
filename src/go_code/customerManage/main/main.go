// 客户关系系统主函数
// Author: 691
// Since: 2023-02-15 16:20
// File: src/go_code/customerManage/quantile/hello.go
package main

import (
	"study/src/go_code/customerManage/service"
	"study/src/go_code/customerManage/view"
)

func main() {
	customerView := view.CustomerView{
		Key:             "",
		Loop:            true,
		CustomerService: service.NewCustomerService(),
	}
	customerView.Mainmenu()
}
