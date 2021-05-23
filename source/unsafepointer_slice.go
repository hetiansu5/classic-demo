package main

import (
	"fmt"
	"unsafe"
)

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func main() {
	// 通过make初始化的slice，如果是0的话，会统一指向一个空数组
	j := make([]int, 0)
	fmt.Println((*sliceHeader)(unsafe.Pointer(&j)))

	// 会初始化一个空元素数组
	var k = []int{}
	fmt.Println((*sliceHeader)(unsafe.Pointer(&k)))

	// 未完成元素数组的初始化
	var z []int
	fmt.Println((*sliceHeader)(unsafe.Pointer(&z)))

	//k变量的地址
	fmt.Println(uintptr(unsafe.Pointer(&k)))
	//k的首元素的值，也就是数组底层结构Data值(指针)
	fmt.Println(*(*uintptr)(unsafe.Pointer(&k)))
}
