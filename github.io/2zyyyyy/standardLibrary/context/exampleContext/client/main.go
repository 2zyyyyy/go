package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// client端
type resData struct {
	res *http.Response
	err error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		// 请求频繁可定义全局的 client 对象并启用长链接 不频繁的使用短连接
		DisableKeepAlives: true}
	client := http.Client{
		Transport: &transport,
	}
	resChan := make(chan *resData, 1)
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/", nil)
	if err != nil {
		fmt.Printf("new request failed, err:%v\n", err)
		return
	}
	// 使用带超时的 ctx 创建一个新的 client request
	req = req.WithContext(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		res, err := client.Do(req)
		fmt.Printf("client.do res:%v, err:%v\n", res, err)
		rd := &resData{
			res: res,
			err: err,
		}
		resChan <- rd
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-resChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(result.res.Body)
		data, _ := ioutil.ReadAll(result.res.Body)
		fmt.Printf("res:%v\n", string(data))
	}
}

func main() {
	// 定义 100ms 超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel() // 调用 cancel 释放 goroutine 资源
	doCall(ctx)
}
