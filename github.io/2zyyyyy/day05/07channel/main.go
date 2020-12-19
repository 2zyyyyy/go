package main

import (
	"fmt"
	"sync"
)

/*channel练习
1.启动一个goroutine,生成100个数发送到ch1
2.启动一个goroutine,从ch1中取值,计算其平方放到ch2中
3.在main中,从ch2取值并打印输出*/

var wg sync.WaitGroup
var once sync.Once

func channelOne(ch1 chan<- int) { // 单向通道：ch1只能接收
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func channelTwo(ch1 <-chan int, ch2 chan<- int) { // ch1只能发送 ch2只能接收
	wg.Done()
	for {
		i, ok := <-ch1 // 通道关闭后再取值ok=false
		if !ok {
			break
		}
		ch2 <- i * i
	}
	once.Do(func() {
		close(ch2) // 确保某个操作只执行一次
	})
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go channelOne(a)
	go channelTwo(a, b)
	go channelTwo(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
