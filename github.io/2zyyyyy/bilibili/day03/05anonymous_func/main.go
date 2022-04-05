package main

import "fmt"

// 匿名函数

func main() {
	// 用变量接收匿名函数
	var f1 = func(x, y int) {
		fmt.Println(x + y)
	}
	f1(1, 99)

	// 如果函数只执调用一次，可以简写成立即执行函数(加一个括号，如有参数括号内传参）
	func(a, b int) {
		fmt.Println("匿名函数立即调用:", a+b)
	}(1, 19)
}
