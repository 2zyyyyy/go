package calc

import "fmt"

func init() {
	fmt.Print("calc: import 时候自动执行！")
}

// 包中的标识符（变量/函数/结构体/接口名等）如果首字母是小写的等价于private
func add(x, y int) int {
	return x + y
}

// 首字母大写的标识符表示对外可见
func Sub(a, b int) int {
	return a - b
}
