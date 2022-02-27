package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 爬虫案例

// 封装错误处理
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 下载图片 传入的是图片叫什么
func DownloadFile(url string, filename string) (ok bool) {
	res, err := http.Get(url)
	HandleError(err, "http.get.url")
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	HandleError(err, "res.body")
	filename = "/Users/gilbert/go/src/go/github.io/2zyyyyy/chineseDocumentation/crawler/image" + filename
	// 写出数据
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

// 并发爬思路
// 1.初始化数据通道
// 2.爬虫写出：26个协程向通道中添加图片链接
// 3.任务统计协程：检查26个任务是否都完成，完成就关闭通道
// 4.下载协程：从通道里读取链接并下载数据

var (
	chanImageUrls chan string // 存放图片链接通道
	wg            sync.WaitGroup
	// 用于监控协程
	chanTask chan string
	reImage  = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {
	// 1.通道初始化
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, 26)
	// 2.爬虫协程
	for i := 0; i < 27; i++ {
		wg.Add(1)
		go get
	}
}

// 下载图片
func DownloadImage() {
	for url := range chanImageUrls {
		filename := G
	}
}

// 获取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

// 任务统计协程
func CheckOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 26 {
			clsoe(chanImageUrls)
			break
		}
	}
	wg.Done()
}

// 爬图片链接存放到通道
// url传的是整页的链接
func getImgUrls(url string) {
	urls := getImgs(url)
	// 遍历切片里所有链接，存入数据管道
	for _, url := range urls {
		chanImageUrls <- url
	}
	// 标识当前协程完成
    // 每完成一个任务，写一条数据
    // 用于监控协程知道已经完成了几个任务
	chanTask <- url
	wg.Done()
}

// 获取当前页面图片链接
func getImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImage)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果/n", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

// 抽取根据url获取内容
func GetPageStr(url string) (page string)  {
	res, err := http.Get(url)
	HandError(err, "http.Get url")
	defer res.
}