package main

import (
	"encoding/json"
	"fmt"
)

// 结构体反序列化
type Hero struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

// 将JSON字符串转换为struct
func unmarshal() {
	str := "{\"Name\":\"张三丰\",\"Age\":98,\"Birthday\":\"2001-09-21\",\"Sal\":3800.85,\"Skill\":\"武当剑法\"}"

	// 定义hero实例
	hero := Hero{}
	err := json.Unmarshal([]byte(str), &hero)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%s\n", err)
	}
	fmt.Printf("反序列化后：%#v\n", hero)
}

func main() {
	unmarshal()
}
