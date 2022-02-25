package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Map

// var m = make(map[string]int)
var m = sync.Map{}

// func get(key string) int {
// 	return m[key]
// }

// func set(key string, value int) {
// 	m[key] = value
// }

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			// strconv.Itoa函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字
			key := strconv.Itoa(n)
			// set(key, n)
			m.Store(key, n)         // 写入
			value, _ := m.Load(key) // 读取
			fmt.Printf("key:%s, value:%d\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
