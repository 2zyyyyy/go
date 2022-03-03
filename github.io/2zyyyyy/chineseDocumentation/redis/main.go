package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Golang 连接 Redis
func main() {
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed, err", err)
		return
	}
	fmt.Println("redis conn success!")

	defer client.Close()

	_, err = client.Do("MSet", "abc", 100, "efg", 200)
	if err != nil {
		fmt.Println("client do failed, err:", err)
		return
	}

	r, err := redis.Ints(client.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	// 遍历r 批量取
	for k, v := range r {
		fmt.Printf("k:%v, v:%v\n", k, v)
	}
}
