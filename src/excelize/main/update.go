package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func updateExcel() {
	billTemplate, err := excelize.OpenFile("Net.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	err = billTemplate.AddChart("账单", "C23", &excelize.Chart{
		Type: excelize.Line,
		Series: []excelize.ChartSeries{
			{
				Name:       "Sheet1!$B$1",
				Categories: "Sheet1!$A$2:$A$8641",
				Values:     "Sheet1!$B$2:$B$8641",
				Line: excelize.ChartLine{
					Smooth: true,
				},
			},
		},
		Legend: excelize.ChartLegend{
			Position:      "top",
			ShowLegendKey: false,
		},
		Dimension: excelize.ChartDimension{
			Width:  1000,
			Height: 400,
		},
		ShowBlanksAs: "zero",
	})
	if err != nil {
		fmt.Println(err)
	}

	err = billTemplate.SaveAs("NewNet.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// 修改表格
	updateExcel()
}
