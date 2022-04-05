package main

import "fmt"

/* 切片:是一个拥有相同类型元素的可变长度的序列
他是基于数组类型做的一层封装，他非常灵活，支持自动扩容
切片是一个引用类型，他的内部结构包含地址、长度和容量。
切片一般用户快速的操作一块数据集合 */

func main() {
	var s1 []int //定义一个存放int类型元素的切片
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"杭州", "金华", "苏州", "嘉兴"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 长度和容量
	fmt.Printf("len(s1)= %d cap(s1)= %d\n", len(s1), cap(s1))
	fmt.Printf("len(s2)= %d cap(s2)= %d\n", len(s2), cap(s2))

	// 由数组获得切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 基于数组切割，[) 左包含右不包含
	fmt.Println("s3=", s3)

	s5 := a1[:4] // [0:4]
	s6 := a1[3:] // [3:len(a1)]
	s7 := a1[:]  // [0:len(a1)]
	fmt.Println(s5, s6, s7)

	// s5 的长度和容量(容量是指底层数组的容量 这里也就是a1的容量)
	fmt.Printf("len(s5): %d cap(s5): %d\n", len(s5), cap(s5))
	// s6 从中间开始切，所以下标前面的容量不计算在s6容量里面
	fmt.Printf("len(s6): %d cap(s6): %d\n", len(s6), cap(s6))

	//切片再切割
	fmt.Println("s6=:", s6)
	s8 := s6[3:]
	fmt.Printf("len(s8): %d cap(s8): %d\n", len(s8), cap(s8))
}
