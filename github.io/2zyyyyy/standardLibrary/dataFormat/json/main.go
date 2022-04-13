package main

import (
	"encoding/json"
	"fmt"
)

// 数据格式

type Equip struct {
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	Occupation   string `json:"occupation"`
	Estate       int64  `json:"estate"`
}

var equip = Equip{
	"破军",
	"北斗第七星，有名破军。破军主破，其利天下无敢拂其锋芒，刀魂暴戾，易走偏锋",
	"东夷战士",
	120,
}

var b = []byte(`{"name":"破军","introduction":"北斗第七星，有名破军。破军主破，其利天下无敢拂其锋芒，刀魂暴戾，易走偏锋","occupation":"东夷战士","estate":120}`)

// 示例通过结构体生成json
func structJson() {
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

// 示例通过map生成json
func mapJson() {
	student := make(map[string]interface{})
	student["name"] = "星河万里"
	student["age"] = 18
	student["sex"] = "man"
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	// 格式化输出
	b, err = json.MarshalIndent(student, "", "	")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))
}

// 示例解析到结构体
func jsonStruct() {
	var e Equip
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
	}
	// 格式化输出
	b, err = json.MarshalIndent(e, "", "	")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))
}

// 示例解析到接口
func jsonInterface() {
	// 声明接口
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(err)
	}
	// 自动转到map
	fmt.Println(i)
	// 可以判断类型
	m := i.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "是float64类型", vv)
		case string:
			fmt.Println(k, "是string类型", vv)
		default:
			fmt.Println("other type", vv)
		}
	}
}

// 面试题： map取一个key，然后修改这个值，原来map数据的值会不会变化
func mapUpdate() {
	student := make(map[string]interface{})
	student["name"] = "星河万里"
	student["age"] = 18
	student["sex"] = "man"
	fmt.Println("before student:", student)
	// 修改key对应的值
	student["name"] = "万里"
	fmt.Println("after student:", student)
}

func main() {
	//structJson()
	//mapJson()
	//jsonStruct()
	//jsonInterface()
	mapUpdate()
}
