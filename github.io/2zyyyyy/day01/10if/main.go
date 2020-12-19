package main

import "fmt"

// if条件判断

func main() {
	age := 18
	if age >= 18 {
		fmt.Println("大于18，条件1通过！")
	} else {
		fmt.Println("小于18，条件1未通过，通过条件2")
	}

	// 多个条件
	score := 75
	if score >= 90 {
		fmt.Println("优秀！！！")
	} else if score >= 75 {
		fmt.Println("良好~~")
	} else if score >= 60 {
		fmt.Println("及格...")
	} else {
		fmt.Println("不及格???")
	}

	// 特殊写法 作用域
	if num := 9; num >= 0 { // num 局部变量
		fmt.Println("大于等于9")
	} else {
		fmt.Println("小于等于9")
	}
	// fmt.Println(num) num无法找到  这样写节省内存占用
}
