package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// 客户端超时取消示例

// server 随机出现慢响应
func indexHandler(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(2)
	if num == 0 {
		time.Sleep(time.Second * 10) // 耗时10秒的慢响应
		fmt.Fprintf(w, "slow response!")
		return
	}
	fmt.Fprintf(w, "normal response ：）")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
