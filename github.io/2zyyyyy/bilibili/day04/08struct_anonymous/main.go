package main

import "fmt"

// 结构体匿名字段
/* 适用场景
1、字段较少较简单
2、不常用 */
type person struct {
	string
	int
}

func main() {
	// 03-1:结构体匿名字段
	p1 := person{
		"五帝",
		18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
