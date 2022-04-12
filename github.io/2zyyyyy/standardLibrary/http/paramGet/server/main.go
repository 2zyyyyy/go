package main

import (
	"fmt"
	"io"
	"net/http"
)

// server 端
func getHandler(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("r.body.close failed, err:", err)
		}
	}(r.Body)
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status" : "ok"}`
	_, _ = w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/golang/page/get", getHandler)
	_ = http.ListenAndServe("127.0.0.1:9090", nil)
}
