package main

import "fmt"

// 结构体模拟实现其他语言中的-继承

//  动物类
type animal struct {
	name string
}

// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动~~\n", a.name)
}

// 狗类
type dog struct {
	feet int8
	animal
}

// 给dog实现一个狗叫的方法
func (d dog) bark() {
	fmt.Printf("%s在叫：汪汪汪~~\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{name: "旺财"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.bark()
	d1.move()
}
