package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*  使用goroutine和channel实现一个计算int64随机数个位数和的程序
1.开启goroutine循环生成int64类型随机数，发送到jobChan
2.开启24个gorourine从jobChan中取出随机数计算个位数的和，并将结果发送到resultChan
3.主goroutine从resultChan取出结果并打印终端输出*/

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

// 定义两个channel用于存放和接收值
var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

// randomNum 随机数生成函数
func randomNum(random chan<- *job) {
	defer wg.Done()
	// 循环生成int64随机数 发送到jobChan
	for {
		n := rand.Int63()
		newJob := &job{
			value: n,
		}
		random <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

// exchangeNum 计算值的方法
func exchangeNum(random <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	// 从jobChan中取出随机数计算个位数的和，并将结果发送到resultChan
	for {
		job := <-random
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		// 将newResult发送到resultChan
		resultChan <- newResult
	}
}

func main() {
	wg.Add(1)
	go randomNum(jobChan)
	// 开启24个goroutine执行exchange
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go exchangeNum(jobChan, resultChan)
	}
	// 取出结果并输出
	for result := range resultChan {
		fmt.Printf("value:%d, sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()
}
