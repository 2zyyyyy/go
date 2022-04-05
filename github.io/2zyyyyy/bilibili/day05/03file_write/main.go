package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容

func fileWrite01() {
	fileObj, err := os.OpenFile("/Users/Tony/go/src/github.io/2zyyyyy/day05/03file_write/file.txt",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err%v", err)
		return
	}
	defer fileObj.Close()
	// write
	_, _ = fileObj.Write([]byte("WRITE：Write FOR 测试\n"))
	// write string
	_, _ = fileObj.WriteString("WRITE STRING: WRITE STRING FOR 测试!\n")
}

func fileWrite02() {
	fileObj, err := os.OpenFile("/Users/Tony/go/src/github.io/2zyyyyy/day05/03file_write/file.txt",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err%v", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	_, _ = wr.WriteString("WRITE: WRITE FOR bufio!!!") // 写到缓存中
	_ = wr.Flush()                                     // 将缓存中内容写入文件
}

func fileWrite03() {
	str := []byte("WRITE: WRITE FOR ioutil")
	err := ioutil.WriteFile("/Users/Tony/go/src/github.io/2zyyyyy/day05/03file_write/file.txt",
		str, 0666)
	if err != nil {
		fmt.Printf("write file failed, err%v", err)
		return
	}
}

func main() {
	//fileWrite01()
	//fileWrite02()
	fileWrite03()
}
