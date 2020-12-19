package main

import (
	"fmt"
	"strings"
)

// 闭包03
/* 判断传入的文件是否是以指定的名称为后缀的
如果是则返回当前文件名称
不是就返回文件名+判断的后缀
*/
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))     // test.jpg
	fmt.Println(jpgFunc("图片01.jpg")) // test.jpg
	fmt.Println(txtFunc("test"))     // test.txt
	fmt.Println(txtFunc("文档01.txt")) // test.txt
}
