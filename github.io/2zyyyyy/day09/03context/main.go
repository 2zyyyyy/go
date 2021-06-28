package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func example(ctx context.Context) {
	defer wg.Done()
	go example2(ctx)
LOOP:
	for {
		fmt.Println("go context test!(0)")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}

func example2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("go context test!(1)")
		time.Sleep(time.Millisecond * 500)
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
	go example(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutime结束
	wg.Wait()
	fmt.Println("over!")
}
