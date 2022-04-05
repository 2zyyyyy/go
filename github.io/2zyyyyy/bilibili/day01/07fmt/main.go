package main

import "fmt"

// fmt占位符
func main() {
	var n = 100
	// 查看类型
	/* fmt.Printf("%T\n", n) // 类型
	fmt.Printf("%v\n", n) //值
	fmt.Printf("%d\n", n) //十进制
	fmt.Printf("%b\n", n) //二进制
	fmt.Printf("%o\n", n) //八进制
	fmt.Printf("%x\n", n) //十六进制
	fmt.Printf("%c\n", n) //字符
	fmt.Printf("%s\n", n) //字符串
	fmt.Printf("%p\n", n) //指针
	fmt.Printf("%v\n", n) //值
	fmt.Printf("%f\n", n) //浮点数
	fmt.Printf("%t\n", n) //布尔值

	s := "golong学习第一天"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s) */
	percentage(n)
	// 整数->字符
	fmt.Printf("%q\n", 97)

	//  字符串
	fmt.Printf("%q\n", "精绝古城")

	// 浮点数和复数
	fmt.Printf("%b\n", 3.14159263873124)
}

// 格式化输出百分比
func percentage(x int) {
	fmt.Printf("%d%%\n", x)
}
