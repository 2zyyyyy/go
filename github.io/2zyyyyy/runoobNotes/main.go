package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const MAX int = 3

// 数组
func list() {
	balance := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance)
}

// 多维数组
func multidimensionalArray() {
	// Step 1: 创建数组
	var array [][]int

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

func forPrintArray() {
	/* 数组 - 5 行 2 列*/
	array := [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}

	/* 输出数组元素 */
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, array[i][j])
		}
	}
	// 创建空的二维数组
	var animals [][]string

	// 创建三一维数组，各数组长度不同
	row1 := []string{"fish", "shark", "eel"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	// 循环输出
	for i := range animals {
		fmt.Printf("row=%v, animals=%v\n", i, animals[i])
	}
}

// 向函数传递数组
func getAverage(array []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += array[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}

// 指针数组
func ptrArray() {
	array := [MAX]int{10, 100, 200}
	var ptr [MAX]*int

	// 循环赋值(将array的地址赋值给ptr)
	for i := range array {
		ptr[i] = &array[i]
	}

	for i, x := range ptr {
		fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i, *x, x)
	}
}

// 指向指针的指针
func ptrPtrValue() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	// 指针ptr的地址
	ptr = &a

	// 指向指针pptr的地址
	pptr = &ptr

	// 获取pptr(指向指针的指针的值)
	fmt.Printf("变量a = %d\n", a)
	fmt.Printf("变量*ptr = %d\n", *ptr)
	fmt.Printf("变量**ptr = %d\n", *pptr)
}

// 向函数传递指针参数
func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func swapExample() {
	// 定义局部变量
	var a, b = 100, 200

	swap(&a, &b)
	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

// 指针
func pointerAddress() {
	a := 10
	fmt.Printf("变量a的地址：%x\n", &a) // %x 16进制 小写字母
}

func pointerOutput() {
	var ip *int // 声明指针变量
	a := 10     // 声明实际变量

	ip = &a // 指针变量的存储地址

	fmt.Printf("a变量的地址：%x\n", &a)

	// 指针变量的存储地址
	fmt.Printf("ip变量存储的指针地址：%x\n", ip)

	// 使用指针访问值
	fmt.Printf("*ip变量的值：%d\n", *ip)
}

func pointerNil() {
	var ptr *int
	fmt.Printf("ptr的值为：%x\n", ptr)
}

// Books struct
type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func structBook() {
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
		bookId:  10001,
	}
	fmt.Println("使用K:V格式创建的结构体：", english2)

	// 忽略的字段为0或空
	english3 := Books{
		title:  "'english book unit 1'",
		author: "2zyyyyy",
	}
	fmt.Println(english3)
}

// BooksMember 结构体成员
type BooksMember struct {
	title   string
	author  string
	subject string
	bookId  int
}

func structMember() {
	var Book1 BooksMember /* 声明 Book1 为 BooksMember 类型 */
	var Book2 BooksMember /* 声明 Book2 为 BooksMember 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.bookId)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.bookId)
}

// 结构体作为函数参数
func argumentFuncStruct(books Books) {
	fmt.Printf("Books title:%v\n", books.title)
	fmt.Printf("Books author:%v\n", books.author)
	fmt.Printf("Books subject:%v\n", books.subject)
	fmt.Printf("Books book_id:%v\n", books.bookId)
}

// structPointer
func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}

func structPointer() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700

	/* 打印 Book1 信息 */
	printBook(&Book1)

	/* 打印 Book2 信息 */
	printBook(&Book2)
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

func structJson() {
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

func sliceLenCap() {
	nums := make([]int, 3, 5)
	printSlice(nums)
}

func sliceNil() {
	var nums []int
	printSlice(nums)
}

func sliceSubstring() {
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

	numsOne := make([]int, 0, 5)
	printSlice(numsOne)

	// 打印子切片从索引[0, 2)
	numsTwo := nums[:2]
	printSlice(numsTwo)

	// 打印索引[0, 2)
	numsThree := nums[2:5]
	printSlice(numsThree)
}

func sliceAppendCopy() {
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
	numsCapDouble := make([]int, len(nums), (cap(nums))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numsCapDouble, nums)
	printSlice(numsCapDouble)
}

func sliceCap() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(nums)

	// 切割slice后获取切片的cap
	numsCap := nums[5:8]
	printSlice(numsCap) // len=3 cap=5 slice=[6 7 8] capacity 为 7 是因为 number3 的 ptr 指向第三个元素， 后面还剩 2,3,4,5,6,7,8, 所以 cap=7。
}

// range
func goRange() {
	nums := []int{1, 3, 5, 7, 9}
	sum := 0
	for _, num := range nums {
		fmt.Println(sum)
		sum += num
	}
	fmt.Printf("sum = %d\n", sum)

	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号
	//所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		fmt.Printf("index = %d, num = %d\n", i, num)
	}

	//range也可以用在map的键值对上。
	maps := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range maps {
		fmt.Printf("key = %s, value = %s\n", k, v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	// 通过 range 获取参数列表:
	fmt.Println(len(os.Args))
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}

// map
func mapExample() {
	countryCapitalMap := make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "的首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American的首都是", capital)
	} else {
		fmt.Println("American的首都不存在")
	}
}

// map delete
func mapDelete() {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")
	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")
	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}

// 递归函数-阶乘（n!=(n-1)! * n）
func factorial(n uint64) (res uint64) {
	if n > 0 {
		res = n * factorial(n-1)
		return res
	}
	return 1 // 0! = 1
}

// 斐波那契数列(F(0)=0，F(1)=1, F(n)=F(n - 1)+F(n - 2)（n ≥ 2，n ∈ N*）)
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

// 类型转换
func typeConversion() {
	var sum = 17
	var count = 5

	mean := float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}

// 接口（interface）
/*
	1 定义接口
	2 定义结构体
	3 实现接口方法
		3.1 方法实现
*/
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (NokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (IPhone IPhone) call() {
	fmt.Println("I am IPhone, I can call you!")
}

// 错误处理
type DivErr struct {
	etype int // 错误类型
	v1    int // 记录下出错时的除数、被除数
	v2    int
}

// 实现接口方法 error.Error()
func (divErr DivErr) Error() string {
	if divErr.etype == 0 {
		return "除零错误"
	} else {
		return "未知错误"
	}
}

// 除法
func division(a, b int) (int, *DivErr) {
	if b == 0 {
		// 返回错误信息
		return 0, &DivErr{0, a, b}
	} else {
		// 返回正常结果
		return a / b, nil
	}
}

// go并发
func flashSale(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

// channel
func sum(s []int, ch chan int) {
	sum := 0
	for i, v := range s {
		fmt.Printf("i=%d, v=%d, sum=%d\n", i, v, sum)
		sum += v
	}
	ch <- sum
}

// channel 缓存
func chanCache() {
	// 定义缓冲为2的可以存储整数类型的通道
	chanChe := make(chan int, 2)

	// 因为chanChe是带缓冲的通道，我们可以同时发送两个数据,而不用立刻需要去同步读取数据
	chanChe <- 1
	chanChe <- 2

	// 获取这两个数据
	fmt.Printf("chan1 = %d, chan2 = %d\n", <-chanChe, <-chanChe)
}

// channel close()
func chanClose(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	list()
	pointerAddress()
	pointerOutput()
	pointerNil()
	structBook()
	structMember()

	books3 := Books{
		title:   "'pytest接口自动化测试",
		author:  "2zyyyyy",
		subject: "test",
		bookId:  10086,
	}
	argumentFuncStruct(books3)
	structPointer()
	structJson()

	sliceLenCap()
	sliceNil()
	sliceSubstring()
	sliceAppendCopy()
	sliceCap()

	multidimensionalArray()
	forPrintArray()

	/* 数组长度为 5 */
	array := [5]int{1000, 2, 3, 17, 50}

	/* 数组作为参数传递给函数 */
	avg := getAverage(array[:], 5)
	/* 输出返回的平均值 */
	fmt.Printf("平均值为：%f\n", avg)
	ptrArray()

	ptrPtrValue()

	swapExample()

	goRange()

	mapExample()
	mapDelete()

	// 阶乘
	n := 10
	fmt.Printf("%d的阶乘是：%d\n", n, factorial(uint64(n)))
	// 斐波那契数列
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \t", fibonacci(i))
	}

	typeConversion()

	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	// 正确调用
	res, err := division(100, 2)
	if err != nil {
		fmt.Println("(1)failed,", err)
	} else {
		fmt.Println("(1)success, 100/2 = ", res)
	}

	// 错误调用
	res, err = division(100, 0)
	if err != nil {
		fmt.Println("(2)failed, ", err)
	} else {
		fmt.Println("(2)success, 100/0 = ", res)
	}

	// 并发
	go flashSale("(1)goroutine begin~")
	flashSale("(2)goroutine begin~")

	// channel
	s := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	x, y := <-ch, <-ch
	fmt.Printf("x = %d, y = %d, x+y = %d\n", x, y, x+y)

	// channel 缓存区
	chanCache()

	// close
	ch2 := make(chan int, 10)
	fmt.Printf("cap(ch2) = %d\n", cap(ch2))
	go chanClose(cap(ch2), ch2)
	/*range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	会结束，从而在接收第 11 个数据的时候就阻塞了。*/
	for i := range ch2 {
		fmt.Println(i)
	}
}
