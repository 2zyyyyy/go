package main

import (
	"fmt"
	"go/github.io/2zyyyyy/chineseDocumentation/socket_stick/proto"
	"net"
)

// socket_stick/client/main.go
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "今天初七开工，杭州中雪。"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("Encode failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
