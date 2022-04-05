package main

import (
	"fmt"
	"os"
)

// Print
func printDemo() {
	fmt.Print("终端输出该文本信息:")
	name := "月满轩尼诗"
	fmt.Printf("我是%s\n", name)
	fmt.Println("终端单独一行输出内容。")
}

// Fprint
func fprintDemo() {
	_, _ = fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./Fprint.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%s\n", err)
		return
	}
	name := "月满轩尼诗"
	// 向打开的文件句柄中写入内容
	_, _ = fmt.Fprintf(fileObj, "往文件中写入信息：%s", name)
}

// Sprint
func sprintDemo() {
	str1 := fmt.Sprint("月满轩尼诗1")
	name := "月满轩尼诗2"
	age := 18
	str2 := fmt.Sprintf("name:%s, age:%d", name, age)
	str3 := fmt.Sprint("月满轩尼诗3")
	fmt.Println(str1, str2, str3)
}

// 格式化占位符
func formatDemo() {
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", true)
	person := struct {
		name string
		age  int32
	}{"月满轩尼诗", 18}
	fmt.Printf("%v\n", person)
	fmt.Printf("%#v\n", person)
	fmt.Printf("%T\n", person)
	fmt.Printf("100%%\n")
}

// 字符串和[]byte
func byteDemo() {
	s := "月满轩尼诗"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
}

func main() {
	printDemo()
	fprintDemo()
	sprintDemo()
	formatDemo()
	byteDemo()
}
