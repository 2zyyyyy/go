package main

import (
	"fmt"
	"strings"
)

// 字符串   go语言字符串"双引号包裹"， 单引号是字符'你'、'g'
func main() {
	// 单行字符串
	s := "/Users/Tony/go/src/github.io/2zyyyyy"
	fmt.Printf("%#v\n", s)

	// 多行字符串
	s2 := `
		世情薄
			人情恶
				雨送黄昏花易落
	`
	fmt.Println(s2)

	// 字符串相关操作
	fmt.Println(len(s)) // 输出长度

	//字符串拼接
	name := "测试"
	world := "test"

	s3 := name + world // 通过+拼接
	println(s3)

	s4 := fmt.Sprintf("%s%s", name, world) // sprint 返回字符串变量
	fmt.Println(s4)

	// 分隔字符串
	ret := strings.Split(s, "/")
	fmt.Println(ret)

	// 字符串包含
	fmt.Println(strings.Contains(s4, "测试"))

	// 前缀
	fmt.Println(strings.HasPrefix(s4, "测试")) // 返回true
	// 后缀
	fmt.Println(strings.HasSuffix(s4, "测试")) // 返回false

	// 输出字符串出现的位置
	sss := "abc测试字符串@@@c"
	fmt.Println(len(sss))
	fmt.Println(strings.Index(sss, "测试"))
	fmt.Println(strings.LastIndex(sss, "c"))

	// 拼接
	fmt.Println(strings.Join(ret, "*"))
}
