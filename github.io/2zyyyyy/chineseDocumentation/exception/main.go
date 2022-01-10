package main

import (
	"fmt"
)

// go语言之异常

func main() {
	// test()
	// panicChannel(1)
	deferPanic()
}

// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
func deferPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic error!")
	}()

	panic("deferPanic test!")
}

// 向已关闭的通道发送数据引发panic
func panicChannel(n int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ch := make(chan int, 10)
	close(ch)
	ch <- n
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			// fmt.Printf("%T\n", err)
			fmt.Println(err.(string)) // 将 interface{} 转型为具体类型
		}
	}()
	panic("panic error!")
}