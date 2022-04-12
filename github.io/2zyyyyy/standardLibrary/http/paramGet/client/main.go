package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 关于GET请求的参数需要使用Go语言内置的net/url这个标准库来处理。
func main() {
	apiUrl := "http://127.0.0.1:9090/get"
	// url param
	data := url.Values{}
	data.Set("name", "月满轩尼诗")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() // url encode
	fmt.Println(u.String())
	res, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("res body close failed, err:", err)
		}
	}(res.Body)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("readAll failed, err:", err)
		return
	}
	fmt.Println(string(b))
}
