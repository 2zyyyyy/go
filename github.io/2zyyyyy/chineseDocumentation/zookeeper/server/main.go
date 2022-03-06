package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"net"
	"os"
	"time"
)

func main() {
	go startServer("127.0.0.1:8897")
	go startServer("127.0.0.1:8898")
	go startServer("127.0.0.1:8899")

	a := make(chan bool, 1)
	<-a
}

func startServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	fmt.Println(tcpAddr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// 注册zk节点
	// 连接zk
	conn, err := GetConnect()
	if err != nil {
		fmt.Printf("connect zk error:%s\n", err)
	}
	defer conn.Close()

	// zk节点注册
	err = RegisterServer(conn, port)
	if err != nil {
		fmt.Printf("connect zk error:%s\n", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Error:%s\n", err)
			if err != nil {
				return
			}
			continue
		}
		go handleClient(conn, port)
	}
}

func handleClient(conn net.Conn, port string) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("defer close conn failed, err:", err)
		}
	}(conn)

	dayTime := time.Now().String()
	_, _ = conn.Write([]byte(port + ":" + dayTime))
}

func RegisterServer(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func GetConnect() (conn *zk.Conn, err error) {
	zkList := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
