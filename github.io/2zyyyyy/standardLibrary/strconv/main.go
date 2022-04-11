package main

import (
	"fmt"
	"strconv"
)

// Atoi() 字符串转数字（如果传入的字符串参数无法转换为int类型，就会返回错误）
func atoiDemo() {
	str := "100"
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("%s can't convert to int\n", str)
	} else {
		fmt.Printf("type：%T\nvalue：%#v\n", i, i) // type: int  value: 100
	}
}

// Itoa() 数字转字符串
func itoaDemo() {
	i := 200
	str := strconv.Itoa(i)
	fmt.Printf("type:%T\nvalue:%#v\n", str, str) // type:string  value:"200"
}

// Parse系列函数
func parseFunc() {
	// ParseBool
	b, _ := strconv.ParseBool("f")
	fmt.Println("parseBool:", b) // parseBool: false
	// ParseFloat
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println("parseFloat:", f) // parseFloat: 3.1415
	// ParseInt
	i, _ := strconv.ParseInt("-2", 10, 64)
	fmt.Println("parseInt:", i) // parseInt: -2
	// ParseUnit
	u, _ := strconv.ParseUint("-2", 10, 64)
	fmt.Println("parseUint:", u) // parseUint: 0
}

// Format系列函数
func formatFunc() {
	// formatBool
	s1 := strconv.FormatBool(true)
	fmt.Println("formatBool:", s1) // formatBool: true
	// formatBool
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println("formatFloat:", s2) // formatFloat: 3.1415E+00
	// formatBool
	s3 := strconv.FormatInt(-2, 16)
	fmt.Println("formatInt:", s3) // formatInt: -2
	// formatBool
	s4 := strconv.FormatUint(2, 16)
	fmt.Println("formatUint:", s4) // formatUint: 2
}

func main() {
	//atoiDemo()
	//itoaDemo()
	//parseFunc()
	formatFunc()
}
