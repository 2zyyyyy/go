package main

import "fmt"

// 数组
func list() {
	balance := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance)
}

// 指针
func pointer_address() {
	a := 10
	fmt.Printf("变量a的地址：%x\n", &a) // %x 16进制 小写字母
}

func pointer_output() {
	var ip *int // 声明指针变量
	a := 10     // 声明实际变量

	ip = &a // 指针变量的存储地址

	fmt.Printf("a变量的地址：%x\n", &a)

	// 指针变量的存储地址
	fmt.Printf("ip变量存储的指针地址：%x\n", ip)

	// 使用指针访问值
	fmt.Printf("*ip变量的值：%d\n", *ip)
}

func pointer_nil() {
	var ptr *int
	fmt.Printf("ptr的值为：%x\n", ptr)
}

// struct
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func struct_book() {
	// 创建一个新的结构体
	english1 := Books{
		"'english book unit 1'",
		"2zyyyyy",
		"English",
		10001,
	}
	fmt.Println(english1)

	// K:V形式
	english2 := Books{
		title:   "'english book unit 1'",
		author:  "2zyyyyy",
		subject: "English",
		book_id: 10001,
	}
	fmt.Println("使用K:V格式创建的结构体：", english2)

	// 忽略的字段为0或空
	english3 := Books{
		title:  "'english book unit 1'",
		author: "2zyyyyy",
	}
	fmt.Println(english3)

}

func main() {
	list()
	pointer_address()
	pointer_output()
	pointer_nil()
	struct_book()
}
