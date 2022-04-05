package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作

func readFileByRead() {
	// 打开文件
	fileObj, err := os.Open("/Users/Tony/go/src/github.io/2zyyyyy/day05/02open_file/main.go")
	if err != nil {
		fmt.Printf("open file failed, err%v", err)
		return
	}

	//关闭文件
	defer fileObj.Close()

	// 读文件
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err != nil {
			fmt.Printf("read file failed, err%v", err)
			return
		}
		fmt.Printf("读了%d个字节", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

func readFileByBufIo() {
	// 打开文件
	fileObj, err := os.Open("/Users/Tony/go/src/github.io/2zyyyyy/day05/02open_file/main.go")
	if err != nil {
		fmt.Printf("open file failed, err%v", err)
		return
	}

	//关闭文件
	defer fileObj.Close()

	// 创建一个从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed, err%v", err)
			return
		}
		fmt.Print(line)
	}
}

func readFileByIoUtil() {
	ret, err := ioutil.ReadFile("/Users/Tony/go/src/github.io/2zyyyyy/day05/02open_file/main.go")
	if err != nil {
		fmt.Printf("read file failed, err%v", err)
		return
	}
	fmt.Print(string(ret))
}

func main() {
	//readFileByRead()
	//readFileByBufIo()
	readFileByIoUtil()
}
