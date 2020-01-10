package main

import (
	"reflect"
	"fmt"
)



func main() {
	var b interface{}
	b = 1
	t := reflect.TypeOf(b)
	fmt.Println(t.Kind())

	var q *interface{}
	q = &b
	tq := reflect.TypeOf(q)
	vq := reflect.ValueOf(q)
	tq1 := reflect.TypeOf(vq.Elem().Interface())
	fmt.Println(tq.Elem().Kind(), tq1.Kind())
}
