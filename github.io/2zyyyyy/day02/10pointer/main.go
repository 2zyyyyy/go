package main

import "fmt"

// golang 指针(pointer)

func main() {
	// 1、'&':取地址
	// 2、'':根据地址取值
	n := 10
	fmt.Println(&n) // 根据值取内存地址

	address := &n
	fmt.Printf("%T\n", address)

	m := *address
	fmt.Println(m) // 根据内存地址取值

	// var a *int  nil pointer
	var a = new(int)
	fmt.Println(a)
	*a = 100
	fmt.Println(*a)

	/* make和new的区别
	1、make和new都是用来申请内存的
	2、new很少用，一般用来给基本数据类型申请内存，string/int ...，返回的是对应类型的指针(*string、*int)
	3、make是用来给slice、map、chan 申请内存的，make函数返回的是对应的这三个类型本身 */
}
