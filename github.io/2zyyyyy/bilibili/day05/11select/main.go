package main

import "fmt"

// select

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case n := <-ch:
			fmt.Println(n)
		case ch <- i:
		}
	}
}
