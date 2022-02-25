package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 比较互斥锁和原子操作性能

var (
	x    int64
	lock sync.Mutex
	wg   sync.WaitGroup
)

// 普通加锁版本
func add() {
	x++
	wg.Done()
}

// 加强版互斥锁
func mutexAdd() {
	lock.Lock()
	x++
	lock.Unlock()
	wg.Done()
}

// 终极版原子操作
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//  普通 不是并发安全
		// go add()
		// 加强 是并发安全 但性能开销大
		// go mutexAdd()
		// 终极 并发安全 性能优于加锁版
		go atomicAdd()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}
