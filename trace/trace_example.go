package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"runtime/trace"
	"strconv"
	"sync"
	"time"
)


var wg sync.WaitGroup
var httpClient = &http.Client{Timeout: 30 * time.Second}

func SleepSomeTime() time.Duration{
	return time.Microsecond * time.Duration(rand.Int()%1000)
}

func create(readChan chan int) {
	defer wg.Done()
	for i := 0; i < 500; i++ {
		readChan <- getBodySize()
		SleepSomeTime()
	}
	close(readChan)
}

func convert(readChan chan int, output chan string) {
	defer wg.Done()
	for readChan := range readChan {
		output <- strconv.Itoa(readChan)
		SleepSomeTime()
	}
	close(output)
}

func outputStr(output chan string) {
	defer wg.Done()
	for _ = range output {
		// do nothing
		SleepSomeTime()
	}
}

// 获取taobao 页面大小
func getBodySize() int {
	resp, _ := httpClient.Get("https://taobao.com")
	res, _ := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return len(res)
}

func run() {
	readChan, output := make(chan int), make(chan string)
	wg.Add(3)
	go create(readChan)
	go convert(readChan, output)
	go outputStr(output)
}

/**
 * 编译
 * go build trace_example.go
 * ./trace_example
 *
 * 启动trace方式1：端口和
 * go tool trace -http=":8000" trace_example trace.out
 *
 * 启动trace方式2
 * go tool trace trace.out
 *
 * web访问8000端口，查看页面
 * http://localhost:8000
 *
 * More to see https://mp.weixin.qq.com/s/2tUkiKOiD8kSKE6wty1Nhw
 * 	View trace：查看追踪 (按照时间分段，上面我的例子时间比较短，所以没有分段)
	Goroutine analysis：Goroutine 分析
	Network blocking profile：网络阻塞概况
	Synchronization blocking profile：同步阻塞概况
	Syscall blocking profile：系统调用阻塞概况
	Scheduler latency profile：调度延迟概况
	User defined tasks：用户自定义任务
	User defined regions：用户自定义区域
	Minimum mutator utilization：最低 Mutator 利用率
 */
func main() {

	// 讲trace数据输出到trace.out
	f, _ := os.Create("trace.out")
	defer f.Close()

	// trace 的开启和停止
	_ = trace.Start(f)
	defer trace.Stop()
	run()
	wg.Wait()
}
