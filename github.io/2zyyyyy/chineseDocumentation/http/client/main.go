package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//resp, _ := http.Get("http://www.baidu.com")
	//fmt.Println(resp)
	res, err := http.Get("http://127.0.0.1:8000/golang")
	if err != nil {
		fmt.Println("get failed, err:", err)
	}
	defer res.Body.Close()
	// 200 OK
	fmt.Println(res.Status)
	fmt.Println(res.Header)

	buf := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := res.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
