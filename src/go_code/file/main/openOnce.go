// 将文件一次性读取到内存中
// Author: 691
// Since: 2023-03-04 11:33
// File: src/go_code/file/quantile/openOnce.go
package main

import (
	"fmt"
	"os"
)

// 这种适合文件不太大的情况
func main() {
	file := "/Users/691/GolandProjects/test.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	fmt.Println(content) //[]byte
	// 不需要显示关闭
	fmt.Println(string(content))
}
