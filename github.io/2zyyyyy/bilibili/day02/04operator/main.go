package main

import "fmt"

// 运算符

func main() {
	var (
		a = 5
		b = 2
	)
	// 算术运算符：加减乘除 余数
	fmt.Println("a+b=", a+b)
	fmt.Println("a-b=", a-b)
	fmt.Println("a*b=", a*b)
	fmt.Println("a/b=", a/b)
	fmt.Println("a'%'b=", a%b)
	// 自增/自减
	a++
	b--

	// 关系运算符
	fmt.Println(a == b) // golang 是强类型语言 相同类型的变量才能比较
	fmt.Println(a != b) // 不等于
	fmt.Println(a > b)  //大于
	fmt.Println(a >= b) //大小于等于
	fmt.Println(a < b)  // 小于
	fmt.Println(a <= b) // 小于

	// 逻辑运算符
	// 如果分数大于60分而且小于100分
	score := 75
	if score > 60 && score < 100 {
		fmt.Println("好好学习~")
	} else {
		fmt.Println("不用学习...")
	}
	// 如果年龄小于18岁或者年龄大于60岁
	age := 26
	if age > 60 || age < 18 {
		fmt.Println("国家帮养老~~~")
	} else {
		fmt.Println("打工仔...")
	}

	// not 取反
	isMarried := false
	fmt.Println(isMarried)  // false
	fmt.Println(!isMarried) // true

	// 位运算符：针对二进制数
	// 5 = 101  2 = 010 （二进制）
	// &:a按位与
	fmt.Println(5 & 2) // 全1为1 000

	// |：按位或
	fmt.Println(5 | 2) /// 有1为1 111

	// ^:按位异或（不一样则为1）
	fmt.Println(5 ^ 2)

	// <<:将二进制位往左指定位数  5->00000101   00010100
	fmt.Println(5 << 2)

	// >>:将二进制位往右指定位数  5->00000101   00000001
	fmt.Println(5 >> 2)

	// 赋值运算符  给变量赋值
	var x = 10
	x++    // x = x +1
	x--    // x = x -1
	x *= 2 // x = x * 2
	x /= 2 // x = x / 2
	x %= 2 // x = x % 2

	x <<= 2 // x = x << 2  左移
	x &= 2  // x = x & 2   位与
	x |= 2  // x = x | 2   位或
	x ^= 4  // x = x | 2   异或
	x >>= 2 // x = x << 2  右移
}
