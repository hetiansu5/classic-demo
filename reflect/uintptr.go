package main

import (
	"reflect"
	"fmt"
)

func main()  {
	var c uintptr
	rt := reflect.TypeOf(&c)
	rv := reflect.ValueOf(&c)
	fmt.Println(rt.Elem().Kind(), rv.Elem().Kind(), rv.Elem().CanSet())

	rv.Elem().Set(reflect.ValueOf(uintptr(222)))

	fmt.Println(rv)
}
