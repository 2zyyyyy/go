package main

import "fmt"

// for 循环

func main() {
	// 基本格式
	/* for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// 套路1  省略初始语句
	var i = 0
	for ; i <= 3; i++ {
		fmt.Println(i)
	}
	// 套路2 省略结束语句
	var n = 5
	for n < 10 {
		fmt.Println(n)
		n++
	}
	// 无限循环
	for {
		fmt.Println("无限循环")
	}
	// for range循环
	s := "Hi月满轩尼诗"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	} */

	// 哑元变量 不想用到的都直接丢给他
	s := "月满轩尼诗"
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}

	// 练习打印九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Println("")
	}
}
