package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	billTemplate, err := excelize.OpenFile("./src/excelize/Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := billTemplate.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows[len(rows)-1][len(rows[0])-2])

	fmt.Println(rows)
}
