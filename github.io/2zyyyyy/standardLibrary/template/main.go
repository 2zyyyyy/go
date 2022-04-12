package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// HTTP server端代码
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmp, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入 w
	_ = tmp.Execute(w, "www.google.com")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed, err:", err)
		return
	}
}
