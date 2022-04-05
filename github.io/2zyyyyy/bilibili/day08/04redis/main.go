package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initRedis() (err error) {
	fmt.Println("golang连接redis")

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// set/get
func redisExampleSimple() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}

}

// zset
func redisExample() {

}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("conn_pool connection failed, error:%v\n", err)
		return
	}
	fmt.Println("conn_pool connection successed!")

	redisExampleSimple()
}
