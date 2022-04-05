package main

import "fmt"

// 递归 自调用自己 递归一定要有一定的退出条件，适合处理问题相同，问题越来越小的场景
// 阶乘 4！=4*3*2*1

// 计算阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 上台阶面试题 n个台阶 一次可以走一步 也可以走两步 有多少种走法
func jumpStep(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return (n - 1) + (n - 2)
}

func main() {
	ret := f(5)
	fmt.Println(ret)

	n := jumpStep(3)
	fmt.Println(n)
}
