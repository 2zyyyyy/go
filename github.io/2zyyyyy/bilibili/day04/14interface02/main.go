package main

import "fmt"

// 接口的实现
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步~~~")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s！！！\n", food)
}

func (c chicken) move() {
	fmt.Println("鸡冻！")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s！\n", food)
}

func main() {
	var a1 animal // 定义一个接口类型的变量
	fmt.Printf("接口类型:%T\n", a1)

	bc := cat{ // 定义一个cat类型的变量
		name: "西西",
		feet: 4,
	}
	a1 = bc
	fmt.Println(a1)

	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	kfc.eat("牛肉！")
	fmt.Printf("kfc类型:%T\n a1类型:%T\n", kfc, a1)
}
