package main

import (
	"fmt"
)

func functionDemo(x, y int, s string) (int, string) {
	// 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func test(fn func() int) int {
	return fn()
}

// 定义函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

// 定义相互交换值的函数
func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func myfunc(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprint(s, x)
}

// 裸 返回
func add(a, b int) (sum int) {
	sum = a + b
	return
}

// 多返回值可直接作为其它函数调用实参。
func test2() (int, int) {
	return 1, 2
}

func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}

	return x
}

func add1(x, y int) (z int) {
	{ // 不能在一个级别，引发 "z redeclared in this block" 错误。
		var z = x + y
		// return   // Error: z is shadowed during return
		return z // 必须显式返回。
	}
}

// 命名返回参数允许 defer 延迟调用通过闭包读取和修改
func deferReturn(x, y int) (z int) {
	defer func() {
		fmt.Printf("defer z=%d\n", z)
		z += 100
	}()
	z = x + y
	return
}

func main() {
	functionDemo(1, 2, "测试functionDemo函数~")
	// fmt.Printf("n:%d\nstr:%s\n", n, str)

	s1 := test(func() int { // 直接将匿名函数当参数
		return 100
	})
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)
	fmt.Println(s1, s2)

	a, b := 1, 50
	swap(&a, &b)
	fmt.Printf("a=%v, b=%v\n", a, b)

	// 接收可变参数
	// fmt.Println(myfunc("sum:", 1, 100, 1000, 10000))

	//使用slice对象做变参时，必须展开。（slice...）
	nums := []int{1, 2, 3, 4, 5}
	res := myfunc("sum:", nums...)
	fmt.Println(res)

	// 直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
	// sum := add(a, b)
	// fmt.Println(a, b, sum)

	// 多返回值可直接作为其它函数调用实参。
	// fmt.Println(add(test2()))
	// fmt.Println(sum(test2()))

	// fmt.Println(deferReturn(1, 2))

	// 匿名函数
	// getSqrt := func(a float64) float64 {
	// 	return math.Sqrt(a)
	// }
	// fmt.Println(getSqrt(4))

	// function variable
	fn := func() {
		fmt.Println("hello world.")
	}
	fn()

	// function cllection
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	fmt.Println(fns[0](100))
	fmt.Println(fns[1](100))

	// function as field
	d := struct {
		fn func() string
	}{
		fn: func() string { return "hello world!" },
	}
	fmt.Println(d.fn())

	// channel of function
	fc := make(chan func() string, 2)
	fc <- func() string { return "hello world." }
	fmt.Println((<-fc)())
}
