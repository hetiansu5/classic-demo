package main

import (
	"fmt"
	"time"
)

var sum int

// go run -race norace.go 利用 -race 来使编译器报告数据竞争问题
// go run norace.go 不会报竞争错误
func main() {
	go add()
	go add()
	time.Sleep(time.Millisecond)
}

//go:norace
//norace 的作用是：跳过竞态检测。
//使用 norace 除了减少编译时间，我想不到有其他的优点了。
//缺点却很明显，那就是数据竞争会导致程序的不确定性。
func add() {
	sum++
	fmt.Println(sum)
}
