package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello goroutine!", i)
}

func helloGoroutine() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

func main() {
	// helloGoroutine()
	// 主协程退出其他任务是否还会执行
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
