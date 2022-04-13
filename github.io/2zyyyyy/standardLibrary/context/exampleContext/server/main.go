package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// server 端 随机出现慢响应
func indexHandle(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(2)
	if num == 0 {
		time.Sleep(time.Second * 10) // 耗时 10s 的慢响应
		_, _ = fmt.Fprintf(w, "slow response!")
		return
	}
	_, _ = fmt.Fprintf(w, "quick respinse!")
}

func main() {
	http.HandleFunc("/", indexHandle)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
