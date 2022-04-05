package main

import "fmt"

/* 闭包概念：
闭包 = 函数 + 外部变量的引用
底层原理：
1、函数可以作为返回值
2、函数内部查找变量的顺序，现在自己内部找，找不到往外找
*/

// 闭包01

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包02  要求f1(f2) 曲线救国

func f1(f func()) { // 函数做参数
	fmt.Println("this func is f1~~")
	f()
}

func f2(x, y int) {
	fmt.Println("this func is f2~~~\n闭包02：", x+y)
}

// 定义函数f3对函数f2包装
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		// fmt.Println(x) // 现在自己里面找x，在往上找（test函数）
		f(x, y) // f(x,y) 是f3的三个参数（f()、x、y）
	}
	return tmp
}

func main() {
	/* a1 := adder(100)
	a2 := a1(50)
	fmt.Println("闭包01：", a2) */

	ret := f3(f2, 1, 999) // 把原来需要传递2个int类型的参数包装秤一个不需要传参的函数
	// ret()
	f1(ret)
	// f1(f3(f2, 1, 99))
}
