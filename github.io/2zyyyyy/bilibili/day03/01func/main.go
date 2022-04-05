package main

import "fmt"

// 函数：一段代码的封装

func main() {
	f1()
	f2("golang!!!")
	ret := f3(1, 99)
	fmt.Println(ret)
	fmt.Println(f4(5, 1))
}

func f1() {
	fmt.Println("Hello Golang~~~")
}

func f2(name string) {
	fmt.Println("Hello", name)
}

func f3(x, y int) (sum int) {
	sum = x + y
	return
}

func f4(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}
