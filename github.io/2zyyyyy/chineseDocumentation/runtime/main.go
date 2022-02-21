package main

import (
	"fmt"
	"runtime"
)

// runtime

func main() {
	// runtime.Gosched()
	// go func(s string) {
	// 	for i := 0; i < 2; i++ {
	// 		fmt.Println(s)
	// 	}
	// }("runtime.Gosched~")

	// // 主协程
	// for i := 0; i < 2; i++ {
	// 	// 切一下 再次分配任务
	// 	runtime.Gosched()
	// 	fmt.Println("主协程")
	// }

	// runtime.Goexit()
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}
