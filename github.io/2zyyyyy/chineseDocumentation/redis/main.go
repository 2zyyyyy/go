package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Golang 连接 Redis
func main() {
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn conn_pool failed, err", err)
		return
	}
	fmt.Println("conn_pool conn success!")

	defer client.Close()

	// 设置过期时间
	_, err = client.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 批量写入
	//_, err = client.Do("MSet", "abc", 100, "efg", 200)
	//if err != nil {
	//	fmt.Println("client do failed, err:", err)
	//	return
	//}

	// 批量获取
	//r, err := conn_pool.Ints(client.Do("MGet", "abc", "efg"))
	//if err != nil {
	//	fmt.Println("get failed, err:", err)
	//	return
	//}

	// 遍历r 批量取
	//for k, v := range r {
	//	fmt.Printf("k:%v, v:%v\n", k, v)
	//}

	// List 队列操作
	//_, err = client.Do("lpush", "book_list", "abc", "ceg", 300)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//r, err := conn_pool.String(client.Do("lpop", "book_list"))
	//if err != nil {
	//	fmt.Println("get abc failed,", err)
	//	return
	//}

	//fmt.Println(r)

	// hash表
	_, err = client.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(client.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}
