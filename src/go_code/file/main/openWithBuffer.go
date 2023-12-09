// 带缓冲区读文件
// Author: 691
// Since: 2023-03-04 10:09
// File: src/go_code/file/quantile/openWithBuffer.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/691/GolandProjects/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}

	// 及时关闭自愿，否则会有内存泄漏
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("close file err = ", err)
		}
	}(file)

	// 默认缓冲区4096，读取文件时不是一次性的，而是读取一部分，处理一部分，可以处理一些比较大的文件
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		// 输出内容
		fmt.Print(str)
		if err == io.EOF { // 表示文件末尾
			break
		}
	}

	fmt.Println("文件读取结束")
}
