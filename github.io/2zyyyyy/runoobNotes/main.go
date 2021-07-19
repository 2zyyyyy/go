package main

import (
	"encoding/json"
	"fmt"
	"time"
)

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

// 结构体成员
type BooksMember struct {
	title   string
	author  string
	subject string
	book_id int
}

func struct_member() {
	var Book1 BooksMember /* 声明 Book1 为 BooksMember 类型 */
	var Book2 BooksMember /* 声明 Book2 为 BooksMember 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}

// 结构体作为函数参数
func argument_func_struct(books Books) {
	fmt.Printf("Books title:%v\n", books.title)
	fmt.Printf("Books author:%v\n", books.author)
	fmt.Printf("Books subject:%v\n", books.subject)
	fmt.Printf("Books book_id:%v\n", books.book_id)
}

// struct_pointer
func print_book(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func struct_pointer() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	print_book(&Book1)

	/* 打印 Book2 信息 */
	print_book(&Book2)
}

/*
	结构体对象转json
	结构体中属性的首字母大小写问题
	首字母大写相当于 public。
	首字母小写相当于 private。
	注意: 这个 public 和 private 是相对于包（go 文件首行的 package 后面跟的包名）来说的。
	敲黑板，划重点
	当要将结构体对象转换为 JSON 时，对象中的属性首字母必须是大写，才能正常转换为 JSON。
	使用 tag 标记要返回的字段名
*/
type Person struct {
	Age  int    `json:"age"` //标记json名字为age
	Name string `json:"name"`
	Time int64  `json:"-"` // 标记忽略该字段
}

func struct_json() {
	person := Person{
		18,
		"大聪明",
		time.Now().Unix(),
	}
	if result, err := json.Marshal(&person); err == nil {
		fmt.Println(string(result))
	}
}

// 切片（slice）
func slice_len_cap() {
	nums := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums), cap(nums), nums)
}

func slice_nil() {
	var nums []int
	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums), cap(nums), nums)
	if nums == nil {
		fmt.Println("切片是空的")
	}
}

func main() {
	list()
	pointer_address()
	pointer_output()
	pointer_nil()
	struct_book()
	struct_member()

	books3 := Books{
		title:   "'pytest接口自动化测试",
		author:  "2zyyyyy",
		subject: "test",
		book_id: 10086,
	}
	argument_func_struct(books3)
	struct_pointer()
	struct_json()

	slice_len_cap()
	slice_nil()
}
