package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100, // 连接池大小
	})
	ctx, cencel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cencel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Printf("redis 连接成功！%v\n", ctx)
	return nil
}

func Example() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		return
	}
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist!")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func main() {
	initClient()
}
