package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func example() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("go context test!")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go example()
	time.Sleep(time.Second * 3)
	exitChan <- true
	wg.Wait()
}
