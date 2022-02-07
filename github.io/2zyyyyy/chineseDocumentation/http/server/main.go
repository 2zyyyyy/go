package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 单独写回调函数
	http.HandleFunc("/golang", myHandle)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		fmt.Println("listen server failed, err:", err)
		return
	}
}

// handle 函数
func myHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式 get/post/put/delete/update
	fmt.Println("method", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	w.Write([]byte("testInfo!!!"))
}
