package main

import (
	"sync"
	"time"
)

//并行执行任务，等待结束超时机制
func waitGroupTimeout() {
	var wg sync.WaitGroup
	ch := make(chan int, 0)

	wg.Add(2)
	go func() {
		time.Sleep(time.Millisecond * 400)
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Millisecond * 300)
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case <-ch:
		return
	case <-time.After(time.Second):
		return
	}
}
