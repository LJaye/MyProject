// 自己创建一个处理器Handler
// Author: 691
// Since: 2023-05-20 11:05
// File: src/go_web/helloWorld/quantile/main.go
package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {
}

// 需要实现的方法
func (m *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "自己创建的处理器处理请求")
}

func main() {
	myHandler := MyHandler{}
	/*
		通过处理器函数处理请求需要调HandleFunc()
		通过处理器处理请求调Handle()
	*/
	//http.Handle("/myHandler", &myHandler)
	//http.ListenAndServe(":8080", nil)

	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		TLSConfig:   nil,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}
