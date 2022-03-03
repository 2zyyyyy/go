package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// Go etcd (go get go.etcd.io/etcd/clientv3)

// put and get
func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err%v\n", err)
		return
	}
	fmt.Println("connect to etcd success!")
	defer client.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = client.Put(ctx, "lmh", "lmh")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	res, err := client.Get(ctx, "lmh")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range res.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}