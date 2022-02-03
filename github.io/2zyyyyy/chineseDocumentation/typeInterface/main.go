package main

import "fmt"

// 类型与接口

// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。
// 例如：狗可以叫，也可以跑。我们就分别定义Sayer和Runner接口

type Sayer interface {
	say()
}

type Runner interface {
	run()
}

type Dog struct {
	name string
}

// dog 可以同时实现以上两个接口

// 实现Sayer接口
func (d Dog) say() {
	fmt.Printf("%s会叫~\n", d.name)
}

// 实现Runner接口
func (d Dog) run() {
	fmt.Printf("%s会跑~\n", d.name)
}

// 多个类型实现同一接口
// 狗会跑 猫咪也会跑 可以同时实现Runner接口
type Cat struct {
	bread string
}

func (c Cat) run() {
	fmt.Printf("%s也会跑~\n", c.bread)
}

func main() {
	var sayer Sayer
	var runner Runner

	dog := Dog{
		"阿黄",
	}
	sayer, runner = dog, dog
	sayer.say()
	runner.run()

	cat := Cat{
		"加菲猫",
	}
	runner = dog
	runner.run()

	runner = cat
	runner.run()
}
