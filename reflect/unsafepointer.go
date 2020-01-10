package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
	u := uint32(32)
	m := unsafe.Pointer(&u)
	fmt.Println(&u, m)

	v := reflect.ValueOf(m)
	fmt.Println(v.Interface())

}
