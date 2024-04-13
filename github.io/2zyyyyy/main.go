package main

import "fmt"

func test() {
	n := 19
	fmt.Println("19 % 10 =", n%10)
	fmt.Println("19 / 10 =", n/10)
}

func main() {
	n := 199
	sum := 0
	for n > 0 {
		sum += n % 10
		fmt.Println("sum:", sum)
		n = n / 10
		fmt.Println("n:", n)
	}
	fmt.Println("各位数的和sum为:", sum)
	test()

	s := []int{7, 2, 8, -9, 4, 0}
	// [:] :在数字（n）前面则表示从前面开始取数 取n个 反之从后面开始取
	fmt.Println(s[:3]) // 第0个开始取 取三个7 2 8
	fmt.Println(s[3:]) // 取最后三个
	fmt.Println(s[:])
}
