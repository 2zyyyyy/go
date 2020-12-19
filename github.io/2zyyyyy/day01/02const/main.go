package main

// 导入语句
import "fmt"

// 批量声明变量
var (
	name string
	age  int
	sex  bool
)
// 批量声明常量
const (
	n1 = 100
	n2
	n3
)
// iota 枚举
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)
// 题目
const (
	b1 = iota //0
	b2        //1
	_         //2,丢弃的
	b3        // 3
)
// 插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //0+1+1=2
	c4        // 2+1=3
)
// 多个常量声明在一行  在一个const中每新增一行常量声明 iota+1
const (
	d1, d2 = iota + 1, iota + 2 //d1=0+1 d2=0+2
	d3, d4 = iota + 1, iota + 2 //d3=1+1 d4=1+2
)
// 定义数量级
const (
	_  = iota //=0,丢弃
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
// 程序入口
func main() {
	//s输出变量
	name = "月满轩尼诗"
	age = 18
	sex = false
	fmt.Print(sex)
	fmt.Println(age)
	fmt.Printf("name:%s", name)
	fmt.Println()

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}
