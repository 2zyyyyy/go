package main

import "fmt"

// 自定义类型和类型别名

//type后面跟的是类型
type java string     // 自定义类型
type python = string // 别名

func main() {
	var ch java
	ch = "golang"
	fmt.Printf("%T\n", ch)

	var hc python
	ch = "javascripts"
	fmt.Printf("%T\n", hc)

	var r rune
	r = '日'
	fmt.Println(r)
	fmt.Printf("%T\n", r)
}
