package main

import "fmt"

//   多用于函数结束之前结束资源文件、数据库、socket连接
func deferDemo() {
	fmt.Println("start!")
	// defer 把他后面的语句延迟执行，直到函数即将返回的时候
	// 多个defer，后进先出（先进后出） 003 -> 002 -> 001
	defer fmt.Println("defer 001 !!!")
	defer fmt.Println("defer 002 !!!")
	defer fmt.Println("defer 003 !!!")
	fmt.Println("end")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	// deferDemo()
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/* main分析
a = 1  b = 2
1.defer calc("1", 1, calc("10", 1, b))
2.calc("10", 1, 2)  --> "10" 1 2 3
3.defer calc("1", 1, 3)
4.a = 0
5.defer calc("2", 0, calc("20", 0, 2))
6.calc("20", 0, 2) --> "20" 0 2 2
7.defer calc("2", 0, 2)
8.b = 1 b此时用不到 混淆作用
9.calc("2", 0, 2) --> "2" 0 2 2
10.calc("1", 1, 3) --> "1" 1 3 4
输出：
"10" 1 2 3
"20" 0 2 2
"2" 0 2 2
"1" 1 3 4 */
