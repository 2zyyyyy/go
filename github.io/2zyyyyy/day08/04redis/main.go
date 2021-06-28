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
func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx, cencel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cencel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil

}

// ZSet
// func Example() {
// 	zsetKey := "language_rank"
// 	languages := []*redis.Z{
// 		&redis.Z{Score: 90.0, Member: "Golang"},
// 		&redis.Z{Score: 98.0, Member: "Java"},
// 		&redis.Z{Score: 95.0, Member: "Python"},
// 		&redis.Z{Score: 97.0, Member: "JavaScript"},
// 		&redis.Z{Score: 99.0, Member: "C/C++"},
// 	}
// 	// ZADD
// 	num, err := rdb.ZAdd(zsetKey, languages...).Result()
// 	if err != nil {
// 		fmt.Printf("zadd failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("zadd %d succ.\n", num)

// 	// 把Golang的分数加10
// 	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
// 	if err != nil {
// 		fmt.Printf("zincrby failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("Golang's score is %f now.\n", newScore)

// 	// 取分数最高的3个
// 	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
// 	if err != nil {
// 		fmt.Printf("zrevrange failed, err:%v\n", err)
// 		return
// 	}
// 	for _, z := range ret {
// 		fmt.Println(z.Member, z.Score)
// 	}

// 	// 取95~100分的
// 	op := redis.ZRangeBy{
// 		Min: "95",
// 		Max: "100",
// 	}
// 	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
// 	if err != nil {
// 		fmt.Printf("zrangebyscore failed, err:%v\n", err)
// 		return
// 	}
// 	for _, z := range ret {
// 		fmt.Println(z.Member, z.Score)
// 	}
// }

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("连接redis失败，err:%v\n", err)
		return
	}
	fmt.Print("连接redis成功~")

	// zset
	key := "language_rank"
	items := []*redis.Z{
		&redis.Z{Score: 90.0, Member: "Golang"},
		&redis.Z{Score: 98.0, Member: "Java"},
		&redis.Z{Score: 95.0, Member: "Python"},
		&redis.Z{Score: 97.0, Member: "JavaScript"},
		&redis.Z{Score: 99.0, Member: "C/C++"},
	}

	ctx, cencel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cencel()
	// 将元素追加到key
	rdb.ZAdd(ctx, key, items...)
}
