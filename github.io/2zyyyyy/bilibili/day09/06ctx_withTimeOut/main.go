package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context WithTimeOut
// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db conecting...")
		time.Sleep(time.Millisecond * 10) // 假设数据库连接耗时10ms
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	// 设置一个50ms自动超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束

	wg.Wait()
	fmt.Println("over~")
}
