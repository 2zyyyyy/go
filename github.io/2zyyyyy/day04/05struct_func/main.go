package main

import "fmt"

// 构造函数
type person struct {
	name string
	age  int
}

// 构造函数返回值类型：1.结构体（结构体较小） 2.结构体指针（结构体较大，减少程序运行开销）
// 构造函数：约定成俗 new开头
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	t := newPerson("ttt", 19)
	fmt.Println(t)
}
