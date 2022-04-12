package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// net/http get demo
func main() {
	res, err := http.Get("https://2zyyyyy.github.io/")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("defer res.Body close failed, err:", err)
		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed, err:", err)
		return
	}
	fmt.Println(string(body))
}
