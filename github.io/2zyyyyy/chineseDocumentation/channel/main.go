package main

import "fmt"

func receiver(ch chan int) {
	// 接收
	ret := <-ch
	fmt.Println("接收成功~", ret)
}

func main() {
	// ch := make(chan int)
	// // 接收 无缓冲通道 发送前先接收 防止死锁
	// go receiver(ch) // 启用goroutine从通道接收值
	// // sned
	// ch <- 100
	// fmt.Println("发送成功！")

	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}
