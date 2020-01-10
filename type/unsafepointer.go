package main

import (
	"fmt"
	"unsafe"
)

func main() {
	u := uint32(32)
	i := int32(1)
	fmt.Println("u", &u)
	fmt.Println("i", &i)
	p := &i
	fmt.Println(p)
	p = (*int32)(unsafe.Pointer(&u))
	fmt.Println(p)
}
