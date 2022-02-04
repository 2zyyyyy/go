package main

import (
	"fmt"
)

// 类型与接口

// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。
// 例如：狗可以叫，也可以跑。我们就分别定义Sayer和Runner接口

type Sayer interface {
	say()
}

type Runner interface {
	run()
}

// 接口嵌套
type Animal interface {
	Sayer
	Runner
}

type Monkey struct {
	name string
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

// 接口嵌套  让monkey实现Animal
func (m Monkey) say() {
	fmt.Printf("%s会说话~\n", m.name)
}

func (m Monkey) run() {
	fmt.Printf("%s会跑步~\n", m.name)
}

func typeAssert(s interface{}) (info string) {
	info, ok := s.(string)
	if ok {
		fmt.Printf("断言成功, info:%s\n", info)
	} else {
		fmt.Println("断言失败")
	}
	return
}

func typeAssertSwitch(s interface{}) (info string) {
	switch v := s.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("type is unsupport")
	}
	return
}

func main() {
	// var sayer Sayer
	// var runner Runner

	// dog := Dog{
	// 	"阿黄",
	// }
	// sayer, runner = dog, dog
	// sayer.say()
	// runner.run()

	// cat := Cat{
	// 	"加菲猫",
	// }
	// runner = dog
	// runner.run()

	// runner = cat
	// runner.run()

	animal := Monkey{"猴哥"}
	animal.say()
	animal.run()

	// 空接口
	var x interface{}
	str := "2zyyyyy.com"
	x = str
	fmt.Printf("x type = %T, x value = %s\n", x, x)

	num := 100
	x = num
	fmt.Printf("x type = %T, x value = %d\n", x, x)

	bool := true
	x = bool
	fmt.Printf("x type = %T, x value = %v\n", x, x)

	typeAssert("123456789")
	typeAssertSwitch(false)
}
