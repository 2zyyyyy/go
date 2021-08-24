package main

import (
	"go/github.io/2zyyyyy/hello"
	"os"
	//_ "go/github.io/2zyyyyy/hello"
)

// 下划线在代码中
func underline() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/Users/gilbert/go/src/go/golang基础.md")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		_, _ = os.Stdout.Write(buf[:n])
	}
}

func main() {
	// import使用了'_',则会编译报错：./main.go:9:2: undefined: hello
	hello.Hello()
	underline()
}
