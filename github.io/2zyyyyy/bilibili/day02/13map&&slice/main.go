package main

import (
	"fmt"
	"strings"
)

// map与slice组合

func main() {
	// 元素为map类型的切片
	var s1 = make([]map[int]string, 10, 10) // make初始化切片
	// 初始化map
	s1[0] = make(map[int]string, 1)
	s1[0][9999] = "系统异常，请稍后重试！"
	fmt.Println(s1)

	// 值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["杭州"] = []int{1, 2, 3}
	// m1["杭州"] = make([]int, {1,2,3})
	fmt.Println(m1)

	// 统计一个字符串中每个单词出现的次数
	s3 := "what do you do what ?"
	//定义切片并初始化
	count := make(map[string]int)
	// 切割字符串
	s4 := strings.Split(s3, " ")
	for i, v := range s4 {
		fmt.Println(i, v)
		count[v]++ // map值自增
	}
	fmt.Println(count)
}
