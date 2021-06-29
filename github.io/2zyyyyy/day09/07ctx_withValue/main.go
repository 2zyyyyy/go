package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context WithValue
// func WithValue(parent Context, key, val interface{}) Context

type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	TraceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code!")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", TraceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50ms后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worder done~")
	wg.Done()
}

func main() {
	// 设置一个50ms超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	// 通知子goroutine结束
	cancel()
	fmt.Println("over~")
}
