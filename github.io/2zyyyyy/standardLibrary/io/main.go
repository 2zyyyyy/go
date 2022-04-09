package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// io 操作
func ioDemo() {
	var buf [16]byte
	_, _ = os.Stdin.Read(buf[:])
	_, _ = os.Stdin.WriteString(string(buf[:]))
}

// 打开和关闭文件
func openFile() {
	// 只读方式打开当前目录下的 main.go 文件
	file, err := os.Open("./main.go")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("close file failed, err%s\n", err)
			return
		}
	}(file)
	if err != nil {
		fmt.Println("open file failed! err:", err)
		return
	}
	log.Println("文件打开成功~")
}

// 写文件
func writeFile() {
	// 新建文件
	file, err := os.Create("./create.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)
	for i := 0; i < 5; i++ {
		_, _ = file.WriteString("月满\n")
		_, _ = file.Write([]byte("轩尼诗\n"))
	}
}

// 读文件
func readFile() {
	// 打开文件
	file, err := os.Open("./create.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Print(string(content))
}

// 拷贝文件
func copyFile() {
	// 打开文件
	srcFile, err := os.Open("./create.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建新文件
	newFile, err2 := os.Create("./copy.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	// defer 关闭文件
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(srcFile)
	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(newFile)

	// 缓存读取
	buf := make([]byte, 1024)
	for {
		// 从源文件读数据
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕~")
			break
		}
		// 写进去
		lines, err := newFile.Write(buf[:n])
		if err != nil {
			return
		}
		fmt.Printf("写入成功，共写入%d数据\n", lines)
	}
}

func main() {
	//ioDemo()
	//openFile()
	//writeFile()
	//readFile()
	copyFile()
}
