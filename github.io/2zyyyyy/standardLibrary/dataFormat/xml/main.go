package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Server 抽取单个server对象
type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

type Servers struct {
	Name    xml.Name `xml:"servers"`
	Version int      `xml:"version"`
	Servers []Server `xml:"server"`
}

func main() {
	data, err := ioutil.ReadFile("./xml.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 格式化输出
	b, err := json.MarshalIndent(servers, "", "	")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))
}
