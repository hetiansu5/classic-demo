package main

import (
	"reflect"
	"fmt"
)

func main()  {
	var b = 'f'
	t := reflect.TypeOf(b)
	v := reflect.ValueOf(b)

	fmt.Println(t.Kind(), v.Kind(), v.CanSet())
}
