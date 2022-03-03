package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Redis 连接池

var pool *redis.Pool  // 创建Redis连接池

func init() {
	pool = &redis.Pool{
		// 实例化一个连接池
		MaxIdle: 16,  // 初始连接数量
		MaxActive: 0, // redis的最大连接数量（0：不确定）
		IdleTimeout: 300,  // 连接关闭时间300秒（300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			// 要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	client := pool.Get()  // 从连接池，取一个连接
	defer client.Close()  // 函数运行结束，将连接放回连接池

	_, err := client.Do("Set", "wanli", 200)
	if err != nil {
		fmt.Println("redis do failed, err:", err)
		return
	}

	r, err := redis.Int(client.Do("Get", "wanli"))
	if err != nil {
		fmt.Println("get wanli failed, err:", err)
		return
	}
	fmt.Println(r)
	_ = pool.Close() // 关闭连接池
}