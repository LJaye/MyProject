// 拷贝文件、图片、视频等
// Author: 691
// Since: 2023-03-04 14:15
// File: src/go_code/file/quantile/copy.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyFile(srcFilePath string, dstFilePath string) (written int64, err error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Println("open file err = ", err)
		return 0, err
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return 0, err
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func main() {
	srcFilePath := "/Users/691/GolandProjects/flower.jpg"
	dstFilePath := "/Users/691/GolandProjects/copyFlower.jpg"
	_, err := copyFile(srcFilePath, dstFilePath)
	if err != nil {
		fmt.Println("copy file err = ", err)
	} else {
		fmt.Println("拷贝成功")
	}
}
