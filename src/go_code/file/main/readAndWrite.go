// 将一个文件中内容读取，写入到另一个文件中
// Author: 691
// Since: 2023-03-04 13:46
// File: src/go_code/file/quantile/readAndWrite.go
package main

import (
	"fmt"
	"os"
)

func main() {
	filePath1 := "/Users/691/GolandProjects/test.txt"
	filePath2 := "/Users/691/GolandProjects/abc.txt"

	content, err := os.ReadFile(filePath1)
	if err != nil {
		fmt.Println("read file err = ", err)
		return
	}

	err = os.WriteFile(filePath2, content, 0666)
	if err != nil {
		fmt.Println("write file err = ", err)
		return
	}

	// 不需要显示打开和关闭文件（已经封装好了）
}
