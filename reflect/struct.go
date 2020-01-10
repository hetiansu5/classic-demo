package main

import (
	"reflect"
	"fmt"
	"strconv"
)

type info struct {
	Id int `query:"id"`
}

func main() {
	var b interface{}
	b = info{Id:21}
	t := reflect.TypeOf(b)
	v := reflect.ValueOf(b)
	fmt.Println(t.Field(0).Tag, v.Field(0).Kind())

	fmt.Println(strconv.FormatBool(true))

	t.Method(0)
	v.Method(0)

	t.FieldByIndex([]int{1, 2})
	v.FieldByIndex([]int{1, 2})
	t.FieldByName("1")
	t.FieldByNameFunc(func(s string) bool {
		return true
	})
	v.FieldByName("1")
	v.FieldByNameFunc(func(s string) bool {
		return true
	})

	var f = 11.5409400
	fmt.Println(fmt.Sprint(f))
	fmt.Println(strconv.FormatFloat(f, 'f', -1, 32))
}
