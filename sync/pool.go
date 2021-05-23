package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Engine struct {
	Num int
}

func main() {
	var pool sync.Pool

	pool.New = func() interface{} {
		return &Engine{Num: rand.Intn(3)}
	}

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				v := pool.Get().(*Engine)
				time.Sleep(time.Millisecond * 50)
				fmt.Println(v.Num)
				pool.Put(v)
			}
		}()
	}

	time.Sleep(3 * time.Second)
}
