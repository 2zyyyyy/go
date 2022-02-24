package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁

var (
	n  int64
	wg sync.WaitGroup
	// lock   sync.Mutex
	rwLock sync.RWMutex
)

func write() {
	rwLock.Lock() // 加写锁
	n += 1
	time.Sleep(time.Millisecond * 2) // 假设写耗时2毫秒
	rwLock.Unlock()                  // 解写锁
	wg.Done()
}

func read() {
	rwLock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读耗时1毫秒
	rwLock.RUnlock()             // 解读锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
