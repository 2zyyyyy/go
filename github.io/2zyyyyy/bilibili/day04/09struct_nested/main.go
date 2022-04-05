package main

import "fmt"

// 结构体嵌套
type information struct {
	price  int
	weight int
}

type phone struct {
	battery     int
	information // 匿名嵌套结构体
}

type mac struct {
	name string
	info information
}

func main() {
	oneplus7P := phone{
		battery: 4600,
		information: information{
			price:  4499,
			weight: 250,
		},
	}
	fmt.Println(oneplus7P)
	// 现在自己结构体中找这个字段，找不到就去匿名嵌套的结构体中查找该字段
	fmt.Println(oneplus7P.price)
}
