package main

import "fmt"

// 类型断言

func assert1(a interface{})  {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok{
		fmt.Println("断言失败~~")
	}else {
		fmt.Println("传进来的是一个字符串~~", str)
	}
	fmt.Println(str)
}

func assert2(a interface{})  {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("当前传入值为string类型", t)
	case int:
		fmt.Println("当前传入值为int类型", t)
	case bool:
		fmt.Println("当前传入值为bool类型", t)
	case int64:
		fmt.Println("当前传入值为int64类型", t)
	}
}

func main() {
	assert1(1)
	assert2(1)
	assert2(false)
	assert2("字符串")
	assert2(int64(9000))
}