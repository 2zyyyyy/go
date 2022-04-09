package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// bufIo
func writeFile() {
	// w(写) 2 r(读) 4 x(执行) 1
	file, err := os.OpenFile("./bufIo.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Print("open file failed, err:", err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	// 获取write对象
	write := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, err = write.WriteString("月满轩尼诗\n")
		if err != nil {
			return
		}
	}
	// 刷新缓冲区，强制写出
	err = write.Flush()
	if err != nil {
		return
	}
}

func readFile() {
	file, err := os.Open("./bufIo.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
}

func main() {
	writeFile()
	readFile()
}
