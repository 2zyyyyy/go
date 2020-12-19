package main

import "fmt"

// 结构体

// 定义一个人类型
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var m person
	m.name = "2zyyyyy"
	m.age = 18
	m.gender = "男"
	m.hobby = []string{"乒乓球", "篮球", "羽毛球"}
	fmt.Println(m)

	var zl person
	zl.age = 18
	zl.name = "2zzzzzl"
	fmt.Printf("type: %T value: %v \n", zl, zl)

	// 匿名结构体
	var s struct {
		name string
		age  int
	}
	s.age = 18
	s.name = "test"
	fmt.Printf("type: %T value: %v \n", s, s)
}
