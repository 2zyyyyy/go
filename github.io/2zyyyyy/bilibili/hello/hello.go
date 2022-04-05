package hello

import "fmt"

// _ 下划线在import中的应用
func init() {
	fmt.Println("import--init() comme here!")
}

func Hello() {
	fmt.Println("hello~")
}
