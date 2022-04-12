package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context

var wg sync.WaitGroup

//var exit bool

// 基本示例
//func worker() {
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//	}
//	// 如何接收外部命令实现退出
//	wg.Done()
//}
//
//func main() {
//	wg.Add(1)
//	go worker()
//	// 如何优雅的实现结束子goroutine
//	wg.Wait()
//	fmt.Println("over~")
//}

// 全局变量方式

// 全局变量方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易统一
// 2. 如果worker中再启动goroutine，就不太好控制了。
//func worker() {
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//		if exit {
//			break
//		}
//	}
//	wg.Done()
//}
//
//func main() {
//	wg.Add(1)
//	go worker()
//	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
//	exit = true                 // 修改全局变量实现子goroutine的退出
//	wg.Wait()
//	fmt.Println("over")
//}

// 通道方式
// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel
//func worker(exitChan chan struct{}) {
//LOOP:
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//		select {
//		case <-exitChan: // 等待接收上级通知
//			break LOOP
//		default:
//		}
//	}
//	wg.Done()
//}
//
//func main() {
//	var exitChan = make(chan struct{})
//	wg.Add(1)
//	go worker(exitChan)
//	time.Sleep(time.Second * 3)
//	// 给予 goroutine 发送推出的信号
//	exitChan <- struct{}{}
//	close(exitChan)
//	wg.Wait()
//	fmt.Println("over~")
//}

// 官方版
func worker(ctx context.Context) {
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker1")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	// 通知 goroutine 结束
	cancel()
	wg.Wait()
	fmt.Println("over~")
}
