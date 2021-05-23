package main

import (
	"fmt"
	"reflect"
)

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}


func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}

func main() {
	mybar := bar
	Decorator(&mybar, bar)
	mybar("hello,", "world!")


	const nProcess = 5
	var chans [nProcess]<-chan int
	for i, v := range chans {
		fmt.Println(i, reflect.TypeOf(v))
	}

	var chans1 chan int
	chans1 = make(chan int, 2)
	chans1<-1
	chans1<-2
	close(chans1)
	for v := range chans1 {
		fmt.Println(v)
	}
}
