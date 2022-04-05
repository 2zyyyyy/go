package main

import "fmt"

// 接口：接口是一种特殊的类型，他规定了变量有哪些方法

// 引出接口的实例

// 定义一个能说话的方法
type speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker类型
}

type cat struct{}

type dog struct{}

type person struct{}

type monkey struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~~~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~~~")
}

func (p person) speak() {
	fmt.Println("嘤嘤嘤~~~")
}

func (m monkey) junp() {
	fmt.Println("跳一跳~~~")
}

func da(x speaker) {
	// 接收一个参数，传进来什么我就打谁
	x.speak() // 被打了就要说话
}

func main() {
	var p1 person
	var c1 cat
	var d1 dog
	// var m1 monkey

	da(p1)
	da(c1)
	da(d1)
	// da(m1)
}
