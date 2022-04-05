package main

import (
	"bufio"
	"fmt"
	"os"
)

func useScan() {
	var str string
	fmt.Print("请输入内容：")
	_, _ = fmt.Scanln(&str)
	fmt.Printf("你输入的内容是：%s\n", str)
}

func userBufio() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", str)
}

func main() {
	//useScan()
	userBufio()
}
