// 判断一个文件是否存在
// Author: 691
// Since: 2023-03-04 14:12
// File: src/go_code/file/quantile/isExist.go
package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "/Users/691/GolandProjects/test.txt"
	res, err := isExist(filePath)
	if err != nil {
		fmt.Println("error!")
		return
	}
	fmt.Println(res)
}

func isExist(fileName string) (bool, error) {
	_, err := os.Stat(fileName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
