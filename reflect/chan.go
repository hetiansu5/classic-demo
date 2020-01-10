package main

import (
	"reflect"
	"fmt"
)

func main()  {
	var c <- chan int
	c = make(chan int, 2)
	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)
	fmt.Println(v.Kind(), v.Cap(), t.ChanDir(), t.Name())
}
