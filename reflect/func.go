package main

import (
	"reflect"
	"fmt"
)

type c func(name string, nums ...int) string

func main() {
	var b c
	t := reflect.TypeOf(b)
	v := reflect.ValueOf(b)
	fmt.Println(v.Kind(), t.NumIn(), t.NumOut(), t.In(0), t.Out(0), t.In(1), t.IsVariadic())
}
