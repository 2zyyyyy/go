package main

import "fmt"

// Golang 方法

// struct
type User struct {
	Name, Address string
}

// 接收者T和*T
type Data struct {
	x int
}

// 普通函数与方法的区别(在接收者分别为值类型和指针类型的时候)

// 1.普通函数
// 接收值类型参数的函数
func valueTest(a int) int {
	return a + 10
}

// 接收指针类型参数的函数
func pointTest(a *int) int {
	return *a + 10
}

func structTestValue() {
	a := 2
	fmt.Println("valueTest:", valueTest(a))
	// 函数的参数作为值类型，则不能直接将指针作为参数传递
	// fmt.Println("valueTest:", valueTest(&a)) // 错误写法

	b := 5
	fmt.Println("pointerTest:", pointTest(&b))
	// 同样，当函数的参数为指针类型时，也不能直接将值类型作为参数传递
	// fmt.Println("pointTest:", pointTest(b)) // 错误写法
}

// 2.方法
type User2 struct {
	id   int
	name string
}

// 接收者为值类型
func (u User2) valueShowName() {
	fmt.Println(u.name)
}

// 接收者为指针类型
func (u *User2) pointShowName() {
	fmt.Println(u.name)
}

func structTestFunc() {
	// 值类型调用方法
	userValue := User2{1, "张三"}
	userValue.valueShowName()
	userValue.pointShowName()

	// 指针类型调用方法
	userPoint := &User2{2, "李四"}
	userPoint.valueShowName()
	userPoint.pointShowName()
	//与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用
}

// struct metnhod
func (u *User) Express(num string) {
	fmt.Printf("u.num=%s, u.name=%s, u.address=%s\n", num, u.Name, u.Address)
}

func (d Data) vauleTest() {
	fmt.Printf("valueTest=%p\n", &d)
}

func (d *Data) pointerTest() {
	fmt.Printf("pointerTest=%p\n", d)
}

func main() {
	// 值类型调用方法
	// u1 := User{
	// 	"张三",
	// 	"法外狂徒张三的家在哪里？",
	// }
	// fmt.Printf("u1 type=%T\n", u1)
	// u1.Express("88798871")

	// 指针类型调用方法
	// u2 := &User{
	// 	"李四",
	// 	"法外狂徒李四的家在哪里？",
	// }
	// fmt.Printf("u2 type=%T\n", u2)
	// u2.Express("15099893012")

	// d := Data{}
	// p := &d
	// fmt.Printf("&d=%p\n", p) // 0xc0000b4008

	// d.vauleTest()   // 0xc0000b4018
	// d.pointerTest() // 0xc0000b4008

	// p.vauleTest()   // 0xc0000b4030
	// p.pointerTest() // 0xc0000b4008

	structTestValue()
	structTestFunc()
}
