package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// byte 和 rune 类型

// go语言为了处理非ASCII码类型的字符 定义了新的rune类型

func main() {
	s := "月满轩尼诗"
	// len() 求得是byte字节的数量
	n := len(s)
	fmt.Println(n)

	/* for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Printf("%c\n", s[i]) // %c 字符
	} */

	for _, c := range s { // 从字符串中拿出具体的字符
		fmt.Printf("%c\n", c) // %c: 字符
	}

	// 字符串修改 原则上不能修改 需要转换成其他变量
	s2 := "我是被修改的字符串"
	s3 := []rune(s2) // 把字符串强制转换成rune切片
	s3[0] = '你'      // s2是字符串 单独修改第一个需要是字符 故 '你' 单引号包裹
	fmt.Println(string(s3))

	a1 := "绿" // string
	a2 := '绿' //rune(int32)
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

	// 练习：统计一个字符串中汉字的数量
	str := "1234我是汉字I'm man!@#$%^&*()_+"
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println("当前字符串：" + str + "\n" + "共有：" + strconv.Itoa(count) + "个中文汉字")
}
