package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"io/ioutil"
	"net"
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
		time.Sleep(1 * time.Second)
	}
}

func startClient() {
	// service := "127.0.0.1:8899"
	// 获取地址
	serverHost, err := getServerHost()
	if err != nil {
		fmt.Printf("get server host fail: %s\n", err)
		return
	}
	fmt.Println("connect host:" + serverHost)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverHost)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()

	_,err = conn.Write([]byte("timestamp"))
	checkError(err)

	res, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(res))

	return
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
