package main

import (
	"encoding/json"
	"fmt"
)

// 数据格式

type Equip struct {
	Name         string
	Introduction string
	Occupation   string
	Estate       int64
}

func main() {
	equip := Equip{
		"破军",
		"北斗第七星，有名破军。破军主破，其利天下无敢拂其锋芒，刀魂暴戾，易走偏锋",
		"东夷战士",
		120,
	}
	// 编码 json
	b, err := json.Marshal(equip)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))

	// 格式化输出
	b, err = json.MarshalIndent(equip, "", "	")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))
}
