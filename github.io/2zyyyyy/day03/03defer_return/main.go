package main

import "fmt"

// golang 函数中的return不是原子操作，底层分为2步执行
// 1、返回值赋值  2、真正ret返回
// 如果函数中存在defer，那么defer执行的时机是第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++ // 返回5，修改的是x,而不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回x=6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 修改的是x的值，不影响之前的x赋值给y=5
	}()
	return x // 返回:y = x = 5
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中的副本
	}(x)
	return 5 // 返回值=x=5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
