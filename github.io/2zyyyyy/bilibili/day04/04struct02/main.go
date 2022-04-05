package main

import "fmt"

type person struct {
	name, gender string
}

// go语言函数传参永远是copy
func f(x person) {
	x.gender = "男" // 修改的是副本的gender的值
}

// 取内存地址 修改其值
func f2(x *person) {
	// (*x).gender = "男" // 修改的是副本的gender的值
	x.gender = "男" // go语法糖 自动根据指针找对应变量
}

func main() {
	var p person
	p.name = "test"
	p.gender = "女"
	f(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)

	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Println(&p2)

	var p3 = person{
		name:   "xx",
		gender: "yy",
	}
	fmt.Printf("%#v\n", p3)

	p4 := person{
		"xxxx",
		"yyy",
	}
	fmt.Printf("%#v\n", p4)

	// 结构体占用一块连续的内存
	type test struct {
		a, b, c int8
	}
	m := test{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
}
