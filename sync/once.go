package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
		doOnce()
	}()

	doOnce()
}


func doOnce(){
	fmt.Println("do once")
	//once执行的function如果panic,仍算return过了，即执行过了
	once.Do(func() {
		panic("dddd")
	})
}
