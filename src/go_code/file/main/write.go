// 创建一个新文件，写入内容
// Author: 691
// Since: 2023-03-04 11:07
// File: src/go_code/file/quantile/write.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		const (
		    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
		    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
		    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
		    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
		    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
		    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
		    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
		    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
		)
	*/
	filePath := "/Users/691/GolandProjects/newFile.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	// 准备写入，使用带缓存的Write
	str := "Hello World!\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writer是带缓存的，因此需要刷到磁盘上，否则会丢失数据
	err = writer.Flush()
	if err != nil {
		fmt.Println("flush file err = ", err)
		return
	}
}
