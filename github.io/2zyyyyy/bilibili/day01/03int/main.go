package main

import "fmt"

//整型

func main() {
	// 十进制
	var n1 = 101
	fmt.Printf("%d\n", n1)
	fmt.Printf("%b\n", n1) //10进制->2进制
	fmt.Printf("%o\n", n1) //10进制->8进制
	fmt.Printf("%x\n", n1) //10进制->16进制

	//八进制
	n2 := 077
	fmt.Printf("%d\n", n2)
	fmt.Printf("")

	//十六进制
	n3 := 0x1234567
	fmt.Printf("%d\n", n3)
}
