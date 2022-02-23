package main

import (
	"fmt"
	"time"
)

// channel select

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1 func"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2 func"
}

// 判断通道是否存满
func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write hello~")
		default:
			fmt.Println("channel full!")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// 两个管道
	// out1 := make(chan string)
	// out2 := make(chan string)
	// // 跑2个子协程 写数据
	// go test1(out1)
	// go test2(out2)
	// // 用select监控
	// select {
	// case str1 := <-out1:
	// 	fmt.Println("str1:", str1)
	// case str2 := <-out2:
	// 	fmt.Println("str2:", str2)
	// }

	// int_chan := make(chan int, 1)
	// str_chan := make(chan string, 1)
	// go func() {
	// 	int_chan <- 1
	// }()
	// go func() {
	// 	str_chan <- "test"
	// }()
	// select {
	// case value := <-int_chan:
	// 	fmt.Println("int value=", value)
	// case value := <-str_chan:
	// 	fmt.Println("string value=", value)
	// }
	// fmt.Println("main结束~")

	// 判断通道是否存满
	ch := make(chan string, 10)
	// 子协程写数据
	go write(ch)
	// 取数据
	for s := range ch {
		fmt.Println("res=", s)
		time.Sleep(time.Second)
	}
}
