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
}
