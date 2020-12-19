package main

import "fmt"

// map

func main() {
	var mp map[string]int
	fmt.Println(mp == nil) // 没有初始化(未在内存中开辟空间)
	// make初始化map 估算好map容量，避免动态扩容
	mp = make(map[string]int, 10)
	mp["测试"] = 100
	mp["test"] = 99
	mp["dev"] = 120
	fmt.Println(mp)

	// 获取value
	fmt.Println(mp["test"])
	// 如果接收的key不存在
	value, ok := mp["不存在的key"]
	if !ok {
		fmt.Println("查无此key~~~")
	} else {
		fmt.Println(value)
	}
	// map的遍历
	for k, v := range mp {
		fmt.Println(k, v)
	}
	// 遍历key
	for k := range mp {
		fmt.Println(k)
	}
	// 遍历value
	for _, v := range mp {
		fmt.Println(v)
	}
	// 删除
	delete(mp, "111")
	fmt.Println(mp)
}
