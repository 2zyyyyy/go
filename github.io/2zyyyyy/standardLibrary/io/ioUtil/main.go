package main

import (
	"fmt"
	"io/ioutil"
)

func writeFile() {
	err := ioutil.WriteFile("./ioUtil.txt", []byte("月满轩尼诗"), 0666)
	if err != nil {
		fmt.Println("ioUtil write file failed, err:", err)
		return
	}
}

func readFile() {
	content, err := ioutil.ReadFile("./ioUtil.txt")
	if err != nil {
		fmt.Println("ioUtil read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	writeFile()
	readFile()
}
