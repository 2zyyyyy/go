package main

import "fmt"

// defer 延迟调用

func main() {
	// deferClosure()
	// ts := []Test{
	// 	{"a"},
	// 	{"b"},
	// 	{"c"},
	// }
	// for _, v := range ts {
	// 	defer Close(v)
	// }
	test(0)
}

func test(x int) {
	defer fmt.Println("a")
	defer fmt.Println("b")

	defer func() {
		fmt.Println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()
	defer fmt.Println("c")
}

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func Close(t Test) {
	t.Close()
}

func deferDemo() {
	var whatever [5]struct{}
	for i := range whatever {
		defer fmt.Println(i)
	}
}

func deferClosure() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}
