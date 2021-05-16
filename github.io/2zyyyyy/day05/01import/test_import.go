package main

// 包导入

import (
	"fmt"
	"go/github.io/2zyyyyy/day05/calc"
)

var x = 100

const pi = 3.14

func init() {
	fmt.Print("自动执行")
	fmt.Print(x, pi)
}

func main() {
	res := calc.Sub(20, 10)
	fmt.Print(res)
}
