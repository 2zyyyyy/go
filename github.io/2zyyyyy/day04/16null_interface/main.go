package main

import (
	"fmt"
)

// 空接口

// interface: 关键字
// interface{}: 空接口

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "万里"
	m1["age"] = 10000
	m1["married"] = false
	m1["hobby"] = [...]string{"唱", "跳", "RAP"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}
