// 获取命令行的各个参数
// Author: 691
// Since: 2023-03-04 15:45
// File: src/go_code/file/quantile/flag.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	//fmt.Println(len(os.Args))
	//for i, arg := range os.Args {
	//	fmt.Printf("args[%v] = %v", i, arg)
	//}

	// flag包解析命令行参数
	// 定义变量，用于接受命令行参数值
	var user, pwd, host string
	var port int
	// &user	就是接受用户输入的 -u 后面的参数值
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	flag.Parse()
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v", user, pwd, host, port)

	// 终端执行 go run flag.go -u liu
}
