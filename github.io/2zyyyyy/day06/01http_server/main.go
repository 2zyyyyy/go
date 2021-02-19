package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./html.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	name := data.Get("name")
	age := data.Get("age")
	fmt.Println(name, age)
	answer := `{"status": "GET OK"}`
	w.Write([]byte(answer))
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println("form:", r.PostForm) // 打印form表单数据
	fmt.Println("name:"+r.PostForm.Get("name"), "age:"+r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.body failed, err:%v\n", err)
		return
	}
	fmt.Println("b:", string(b))
	answer := `{"status":"POST OK"}`
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/golang/page", f1)
	http.HandleFunc("/golang/page/get", getRequest)
	http.HandleFunc("/golang/page/post", postRequest)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
