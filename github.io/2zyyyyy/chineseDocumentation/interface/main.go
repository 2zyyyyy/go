package main

import "fmt"

// interface

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type People interface {
	Speak(string) string
}

type Cat struct{}

type Dog struct{}

type Student struct{}

func (c Cat) say() {
	fmt.Println("喵喵喵")
}

func (d Dog) say() {
	fmt.Println("汪汪汪")
}

// 值接收者实现接口
// func (d Dog) move() {
// 	fmt.Println("dog会动~")
// }

// 指针接收者实现接口
func (d *Dog) move() {
	fmt.Println("dog会动~")
}

func (s *Student) Speak(think string) (talk string) {
	if think == "dsb" {
		talk = "大帅币"
	} else {
		talk = "小水币"
	}
	return
}

func main() {
	// 声明Sayer类型变量
	var sayer Sayer
	// sayer.say()  // panic: runtime error: invalid memory address or nil pointer dereference
	// 实例化dog和cat
	cat := Cat{}
	dog := Dog{}

	sayer = cat
	sayer.say() // 喵喵喵

	sayer = dog
	sayer.say() // 汪汪汪

	var move Mover
	// var dogValue = Dog{} // Dog类型
	// move = dogValue      // move 可以接收Dog类型  // 如果move是指针接收者实现的接口 move不可接收Dog类型
	// move.move()

	var dogPointer = &Dog{} // *Dog类型
	move = dogPointer       // move 可以接收*Dog类型
	move.move()

	// 面试题
	var people People = &Student{}
	think := "the"
	fmt.Println(people.Speak(think))
}
