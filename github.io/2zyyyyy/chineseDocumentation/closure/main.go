package main

import "fmt"

// 闭包（Closure）

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func test() func() {
	x := 100
	fmt.Printf("x(%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x(%p) = %d\n", &x, x)
	}
}

func add(base int) func(int) int {
	return func(i int) int {
		base += 1
		return base
	}
}

func test01(base int) (func(int) int, func(int) int) {
	// 定义两个函数并返回
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

// 数字阶乘（递归案例，我调我自己）
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

// 斐波那契数列(Fibonacci)
func fionacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fionacci(i-1) + fionacci(i-2)
}

func main() {
	// 闭包
	c := a()
	c()
	// c()
	// c()

	a() // 不会输出i
	f := test()
	f()

	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))

	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))

	f1, f2 := test01(10)
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))

	// 递归
	fmt.Println(factorial(7))
	for i := 10; i < 15; i++ {
		fmt.Printf("%d\n", fionacci(i))
	}
}
