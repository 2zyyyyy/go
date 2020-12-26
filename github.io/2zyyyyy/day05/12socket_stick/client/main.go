package main

import (
	"fmt"
	"net"
)

// socket_stick 黏包

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "Hello, 树先生!"
		conn.Write([]byte(msg))
	}
}
