package main

import (
	"context"
	"fmt"
	"time"
)

// context WithDeadLine
// func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

func main() {
	d := time.Now().Add(time.Second * 2)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("case 1")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
