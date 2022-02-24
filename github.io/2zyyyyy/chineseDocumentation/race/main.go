package main

import (
	"fmt"
	"sync"
)

// 竞态问题

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 1000; i++ {
		// 加锁
		lock.Lock()
		x += 1
		// 解锁
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
