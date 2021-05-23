package main

import (
"fmt"
"runtime"
	"time"
)

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			runtime.Gosched()
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
	time.Sleep(time.Second)

	// 让出CPU时间片，重新等待安排任务
	runtime.Gosched()

	// 退出当前协程
	runtime.Goexit()

	// 设置当前程序并发时占用的CPU逻辑核心数
	// Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。
	runtime.GOMAXPROCS(1)
}
