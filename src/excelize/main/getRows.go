package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	billTemplate, err := excelize.OpenFile("./src/excelize/KENC-202310.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := billTemplate.GetRows("账号汇总")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
