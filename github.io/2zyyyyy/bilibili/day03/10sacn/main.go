package main

import "fmt"

// Scan 获取用户输入
func main() {
	/* var s string
	fmt.Scan(&s)
	fmt.Println("用户输入的内容是：", s) */

	var (
		name, class string
		age         int
	)
	/* fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println("学生信息为：", name, age, class) */

	fmt.Scanln(&name, &age, &class)
	fmt.Println("学生信息为：", name, age, class)
}
