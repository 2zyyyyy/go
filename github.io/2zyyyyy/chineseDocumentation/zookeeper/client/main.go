package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		startClient()
	}
}

func startClient() {
	// service := "127.0.0.1:8899"
	// 获取地址
	serverHost, err := getServerHost()
}

func getServerHost() (host string, err error) {
	conn, err := GetConnect()
	if err != nil {
		fmt.Printf("connect zk error:%s\n", err)
		return
	}

	defer conn.Close()
	serverList, err := GetServerList()
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
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
