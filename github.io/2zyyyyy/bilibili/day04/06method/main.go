package main

import "fmt"

// 方法 Method
// 标识符：变量名 函数名 类型名 方法名

// go语言如果标识符的首字母是大写的，那么就表示是公共的 对外可见（public 例如：fmt包下面的Print首字母就是大写的）
type dog struct {
	name string
}

// 构造函数
type person struct {
	name string
	age  int
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 使用值接收者
func (p person) newYear() {
	p.age++
}

/* 指针接收者
什么时候用指针接收者：
1、需要修改接收者中的值
2、接收者是拷贝对象比较大的大对象
3、保持一致性，如果有某个方法用了指针接收者，那么其他方法也尽量用指针接收者 */
func (p *person) realNewYear() {
	p.age++
}

// 方法是做用于特定类型的函数
// (d dog) 接受者表示的是调用该方法的具体变量类型，多用类型名首字母小写表示
func (d dog) eat() {
	fmt.Printf("%s:狗吃屎\n", d.name)
}

// 自定义类型添加方法（不能给其他包里面的类型定义方法）
type myString string

func (m myString) out() {
	fmt.Println("myString的out方法~")
}

func main() {
	dg := newDog("tiantian")
	dg.eat()
	p := newPerson("zy", 18)
	println("调用过年前年龄：", p.age) // 18
	// p.wang() p.wang undefined (type person has no field
	p.newYear()
	println("调用过年后年龄：", p.age) // 18

	p.realNewYear()
	fmt.Println("真过年了，年龄：", p.age)

	s := myString("test")
	s.out()
}
