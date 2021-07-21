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

// 多维数组
func multidimensional_array() {
	// Step 1: 创建数组
	array := [][]int{}

	// Step 2: 使用 appped() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	array = append(array, row1)
	array = append(array, row2)

	// Step 3: 打印两行数据
	fmt.Println("row1:", array[0])
	fmt.Println("row2:", array[1])

	// Step 4: 访问第最后一个元素
	fmt.Println(array[1][2])
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
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)
}

func slice_len_cap() {
	nums := make([]int, 3, 5)
	printSlice(nums)
}

func slice_nil() {
	var nums []int
	printSlice(nums)
	if nums == nil {
		fmt.Println("切片是空的")
	}
}

func slice_substring() {
	/* 创建切片 */
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(nums)
	// 打印原始切片
	fmt.Println("nums:", nums)

	// [1, 4]所以1到4(不包含)
	fmt.Println("nums[1:4]:", nums[1:4])

	// 默认下限为0
	fmt.Println("nums[:3]:", nums[:3])

	// 默认上限weilen(nums)
	fmt.Println("nums[4:]:", nums[4:])

	nums_one := make([]int, 0, 5)
	printSlice(nums_one)

	// 打印子切片从索引[0, 2)
	nums_two := nums[:2]
	printSlice(nums_two)

	// 打印索引[0, 2)
	nums_three := nums[2:5]
	printSlice(nums_three)
}

func slice_append_cppy() {
	var nums []int
	printSlice(nums)

	/* 允许追加空切片 */
	nums = append(nums, 0)
	printSlice(nums)

	/* 向切片添加一个元素 */
	nums = append(nums, 1)
	printSlice(nums)

	/* 同时添加多个元素 */
	nums = append(nums, 2, 3)
	printSlice(nums)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	nums_cap_double := make([]int, len(nums), (cap(nums))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(nums_cap_double, nums)
	printSlice(nums_cap_double)
}

func slice_cap() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(nums)

	// 切割slice后获取切片的cap
	nums_cap := nums[5:8]
	printSlice(nums_cap) // len=3 cap=5 slice=[6 7 8] capacity 为 7 是因为 number3 的 ptr 指向第三个元素， 后面还剩 2,3,4,5,6,7,8, 所以 cap=7。
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
	slice_substring()
	slice_append_cppy()
	slice_cap()

	multidimensional_array()
}
