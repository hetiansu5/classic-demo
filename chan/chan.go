package main

import (
	"time"
	"sync"
)

//多chan select超时机制
func multipleChan() {
	ch := make(chan int, 0)
	select {
	case <-ch:
		return
	case <-time.After(time.Second):
		return
	}
}

//构建一个异步通道
func genChan(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//通过通道实现停止消息通知机制
func doneChane(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	for _, c1 := range cs {
		output(c1)
	}

	return out
}
