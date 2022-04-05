package main

import "fmt"

// 浮点数

func main() {
	//math.MaxFloat32 //float32最大值
	f1 := 1.234556789
	fmt.Printf("%T\n", f1) // 默认go语言中的小数都是float64类型

	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) //显示声明float32类型
}
