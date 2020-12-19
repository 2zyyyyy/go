package main

import (
	"fmt"
	"sync"
)

// channel
// 通道操作：关闭 close() 发送数据 ch1 <- 10 接收数据 n := <- ch1

var a []int
var b chan int // 默认nil,需要手动使用make分配内存空间（make：slice、map、chan）
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)
	// 通道的初始化,否则无法使用
	b = make(chan int) // 不带缓冲区的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		n := <-b
		fmt.Println("后台goroutine从同道b中取到了：", n)
	}()
	b <- 10
	fmt.Println("将10发送到了通道b中")
	wg.Wait()
	close(b)
}

func bufChannel() {
	fmt.Println(b)
	// 通道的初始化,否则无法使用
	b = make(chan int, 16) // 带缓冲区的初始化
	b <- 10
	fmt.Println("将10发送到了通道b中")
	n := <-b
	fmt.Println("后台goroutine从同道b中取到了:", n)
	close(b)
}

func main() {
	noBufChannel()
	fmt.Printf("------------------------------------\n")
	bufChannel()
}
