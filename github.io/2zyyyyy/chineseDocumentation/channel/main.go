package main

import "fmt"

func receiver(ch chan int) {
	// 接收
	ret := <-ch
	fmt.Println("接收成功~", ret)
}

func counter(out chan<- int) { // out是只能发送的channel
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) { // out只能发送 in只能接收
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	// ch := make(chan int)
	// // 接收 无缓冲通道 发送前先接收 防止死锁
	// go receiver(ch) // 启用goroutine从通道接收值
	// // sned
	// ch <- 100
	// fmt.Println("发送成功！")

	// ch := make(chan int)
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()

	// for {
	// 	if data, ok := <-ch; ok {
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }
	// fmt.Println("main结束")

	// // channel 练习
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// // 开启goroutine将0~100数据发送到ch1中
	// go func() {
	// 	for i := 0; i <= 100; i++ {
	// 		ch1 <- i
	// 	}
	// 	close(ch1)
	// }()
	// // 开启goroutine从ch1中接收值 并将该值的平方发送到ch2中
	// go func() {
	// 	for {
	// 		i, ok := <-ch1 // 如果通道关闭再取值ok=false
	// 		if !ok {
	// 			break
	// 		}
	// 		ch2 <- i * i
	// 	}
	// 	close(ch2)
	// }()
	// // 在主goroutine中从ch2接收值并打印
	// for i := range ch2 { // 通道关闭后会退出 for range循环
	// 	fmt.Println(i)
	// }
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
