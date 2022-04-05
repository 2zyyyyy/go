package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notify bool

func example() {
	defer wg.Done()
	for {
		fmt.Println("go context test!")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go example()
	time.Sleep(time.Second * 3)
	notify = true
	wg.Wait()
}
