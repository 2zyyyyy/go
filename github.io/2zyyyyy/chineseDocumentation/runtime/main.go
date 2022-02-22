package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime

// 将任务分配到不同的CPU逻辑核心上实现并行的效果
func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

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
	// go func() {
	// 	defer fmt.Println("A.defer")
	// 	func() {
	// 		defer fmt.Println("B.defer")
	// 		// 结束协程
	// 		runtime.Goexit()
	// 		defer fmt.Println("C.defer")
	// 		fmt.Println("B")
	// 	}()
	// 	fmt.Println("A")
	// }()
	// for {
	// }

	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
