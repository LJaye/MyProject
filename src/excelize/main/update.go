package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func updateExcel() {
	billTemplate, err := excelize.OpenFile("./src/excelize/Net.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	detailSheetName := "边缘计算南昌明细"
	dotNum := 8640
	err = billTemplate.AddChart("账单", "C23", &excelize.Chart{
		Type: excelize.Line,
		Series: []excelize.ChartSeries{
			{
				Name:       fmt.Sprintf("%s!$B$1", detailSheetName),
				Categories: fmt.Sprintf("%s!$A$2:$A$%d", detailSheetName, dotNum+1),
				Values:     fmt.Sprintf("%s!$B$2:$B$%d", detailSheetName, dotNum+1),
				Line: excelize.ChartLine{
					Smooth: true,
				},
				Marker: excelize.ChartMarker{
					Symbol: "none",
				},
			},
		},
		Legend: excelize.ChartLegend{
			Position: "top",
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

	i := 11
	nodes := []string{"fzct11", "ncct49", "xtun03", "zsct02"}
	nodesPriceMap := map[string]float64{
		"fzct11": 8.3,
		"ncct49": 7,
		"xtun03": 7,
		"zsct02": 7.6,
	}
	nodesValueMap := map[string]float64{
		"fzct11": 1573.41,
		"ncct49": 3000.00,
		"xtun03": 1155.36,
		"zsct02": 5534.95,
	}
	for _, node := range nodes {
		err = billTemplate.SetCellValue("账单", "I"+strconv.Itoa(i), nodesValueMap[node])
		err = billTemplate.SetCellValue("账单", "J"+strconv.Itoa(i), nodesPriceMap[node])
		i++
	}

	err = billTemplate.UpdateLinkedValue()

	err = billTemplate.SaveAs("./src/excelize/NewNet.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// 修改表格
	updateExcel()
}
