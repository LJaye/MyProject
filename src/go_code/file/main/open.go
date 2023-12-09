// 文件的打开与关闭
// Author: 691
// Since: 2023-03-04 10:28
// File: src/go_code/file/quantile/open.go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/691/GolandProjects/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}

	// 文件是指针类型
	fmt.Println("file = ", file)

	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err = ", err)
	}
}
