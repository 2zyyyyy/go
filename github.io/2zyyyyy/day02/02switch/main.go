package main

import "fmt"

// switch 简化大量判断

func main() {
	// var n = 3
	switch n := 3; n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("异常输入~")
	}

	// 判断奇偶数
	switch n := 7; n {
	case 1, 3, 5, 7:
		fmt.Println("他们是奇数~~")
	case 2, 4, 6, 8:
		fmt.Println("他们是偶数~！")
	default:
		fmt.Println("猜猜我是谁？")
	}
}
