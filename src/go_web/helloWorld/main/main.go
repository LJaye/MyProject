// 一个简单的web服务器
// Author: 691
// Since: 2023-05-18 15:42
// File: src/go_web/helloWorld/main.go
package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world", r.URL.Path)
}

func main() {
	/**
	HandleFunc可以将函数自动转换成处理器
	但是函数参数必须是(w http.ResponseWriter, r *http.Request)
	*/
	http.HandleFunc("/", handler)
	// 创建路由
	http.ListenAndServe(":8080", nil)
	// 访问：http://localhost:8080    http://localhost:8080/hello

	// 创建多路复用器
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", mux)
}
