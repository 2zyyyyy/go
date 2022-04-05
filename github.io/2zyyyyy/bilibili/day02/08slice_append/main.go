package main

// 切片的append()

import "fmt"

func main() {
	s1 := []string{"杭州", "嘉兴", "金华"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	// 调用append()函数必须使用原来的切片变量接收返回值
	// 必须用变量接收append返回值
	s1 = append(s1, "宁波")                                             // append追加元素，原来的底层数组放不下的时候 go底层就会把底层数组换一个
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) // 自动扩容

	s2 := []string{"武汉", "成都", "苏州"}
	s1 = append(s1, s2...) // s2为切片类型无法直接赋值给s1 需要用...表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
