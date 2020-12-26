package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// socket_stick

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	// 开启循环
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break // 跳出循环 如嵌套多个 跳出最近的内循环
		}
		if err != nil {
			fmt.Println("read from client failed, err", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到来自client发送的数据:", recvStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	// 开启循环
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue // 跳出本次循环 下次继续
		}
		go process(conn)
	}
}
