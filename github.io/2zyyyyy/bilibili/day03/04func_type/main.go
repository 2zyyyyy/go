package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("golang!!!")
}

func f2() int {
	return 10
}

// 函数作为参数
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(x, y int) int {
	return x + y
}

// 函数作为返回值
func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		sum := a + b
		return sum
	}
	return ret
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n%T\n", f3, f4)

	f7 := f5(f2)
	fmt.Printf("%T\n", f7)
}
