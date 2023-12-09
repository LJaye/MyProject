// 统计文件中英文、数字、空格和其他字符的数量
// Author: 691
// Since: 2023-03-04 15:42
// File: src/go_code/file/quantile/statistics.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义结构体，用于保存统计的结果
type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	/**
	1、打开文件，创建Reader
	2、每读取一行都进行一次统计
	3、将结果保存到结构体中
	*/

	file, err := os.Open("/Users/691/GolandProjects/data.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		// 遍历str，进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透处理
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v <= '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				// 其他字符为换行符
				count.OtherCount++
			}
		}
		if err == io.EOF {
			break
		}
	}

	fmt.Println(count)
}
