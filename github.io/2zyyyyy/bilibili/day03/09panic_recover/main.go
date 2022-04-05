package main

import "fmt"

// panic(defer一定要在可能引发panic的语句之前定义)和recover（必须搭配defer,少用）

func funcA() {
	fmt.Println("A")
}

func funcB() {
	// 刚刚打开数据库连接
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接！")
	}()
	panic("数据库连接失败......")
	//fmt.Println("B") // 永远不会执行
}

func funcC() {
	fmt.Println("C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
