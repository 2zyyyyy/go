package main

import (
	"fmt"
	"strconv"
)

// Atoi() 字符串转数字（如果传入的字符串参数无法转换为int类型，就会返回错误）
func atoiDemo() {
	str := "100"
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("%s can't convert to int\n", str)
	} else {
		fmt.Printf("type：%T\nvalue：%#v\n", i, i) // type: int  value: 100
	}
}

// Itoa() 数字转字符串
func itoaDemo() {
	i := 200
	str := strconv.Itoa(i)
	fmt.Printf("type:%T\nvalue:%#v\n", str, str) // type:string  value:"200"
}

func main() {
	atoiDemo()
	itoaDemo()
}
