package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// 客户端
type resData struct {
	res *http.Response
	err error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		// 请求频繁可定义全局的client对象并启用长链接
		// 请求不频繁使用短链接
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}
	resChan := make(chan *resData, 1)
	request, err := http.NewRequest("GET", "http://127.0.0.1:8000/", nil)
	if err != nil {
		fmt.Printf("new request failed, err:%v\n", err)
		return
	}
	// 使用带超时的ctx创建一个新的client request
	request = request.WithContext(ctx)

	var wg sync.WaitGroup

	wg.Add(1)
	defer wg.Wait()
	go func() {
		res, err := client.Do(request)
		fmt.Printf("client.do response:%v, err:%v\n", res, err)
		rd := &resData{
			res: res,
			err: err,
		}
		resChan <- rd
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		//transport.CancelRequest(request)
		fmt.Println("call api timeout!!!")
	case result := <-resChan:
		fmt.Println("call api success~")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer result.res.Body.Close()
		data, _ := ioutil.ReadAll(result.res.Body)
		fmt.Printf("res:%v\n", string(data))
	}
}

func main() {
	// 定义一个100毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	// 调用cancel释放子goroutine资源
	defer cancel()
	doCall(ctx)
}
