package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main(){
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	time.Sleep(3 * time.Second)
        
	calc()  	
 
	c := make(chan bool)
	<-c
}

func calc () {
	i := 1
        for j := 0;j < 10000000; j++ {
                i += j
        }
	fmt.Println(i)
}
