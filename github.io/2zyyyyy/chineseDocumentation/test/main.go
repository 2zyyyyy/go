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

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) []student {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	ce = append(ce, student{3, "xiaowang", 56})
	return ce
}

// 将JSON字符串转换为struct
func unmarshal() {
	str := "{\"Name\":\"张三丰\",\"Age\":98,\"Birthday\":\"2001-09-21\",\"Sal\":3800.85,\"Skill\":\"武当剑法\"}"

	// 定义hero实例
	hero := &Hero{}
	err := json.Unmarshal([]byte(str), &hero)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%s\n", err)
	}
	fmt.Printf("反序列化后：%#v\n", hero)
}

func main() {
	// unmarshal()
	ce := []student{
		{1, "xiaoming", 22},
		{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
}
