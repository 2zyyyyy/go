package main

import (
	"fmt"
	"sort"
)

// slice  copy()

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 // 赋值
	// var a3 []int   nil
	var a3 = make([]int, 3, 3)
	copy(a3, a1) // copy
	fmt.Println(a1, a2, a3)
	a1[0] = 10000
	fmt.Println(a1, a2, a3)

	// 切片中删除元素
	// 将a1索引为1的元素（3）删除
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)

	// 切片练习题
	// 1、写出输出值
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	// 2、使用sort包对数组var a = [...]int{3,7,9,8,1}排序
	var l1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(l1[:]) // 对切片进行排序
	fmt.Println(l1)
}
