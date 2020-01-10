package main

import (
	"reflect"
	"fmt"
)

func main() {
	var arr []interface{}
	arr = append(arr, 7)
	arr = append(arr, 8)
	arr = append(arr, 9)
	arr = append(arr, 10)
	arr = append(arr, 11)
	t := reflect.TypeOf(arr)
	v := reflect.ValueOf(arr)
	fmt.Println(t.Kind(), v.Len(), v.Index(3))

	v.Slice(0,1)

}
