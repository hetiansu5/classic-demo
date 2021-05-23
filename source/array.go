package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var iniSet1 = [6]int{2, 4, 5}
	fmt.Println(iniSet1, len(iniSet1), cap(iniSet1))

	//实际上是[3]int，与[6]int，不是同一个类型
	var iniSet2 = [...]int{2, 4, 5}
	fmt.Println(reflect.TypeOf(iniSet2), reflect.ValueOf(iniSet2).Kind())

	// 验证数组的底层结构，就是指向一块连续的内存。
	// 数组的length是在编译期就确定的了，汇编指令的时候，length就是常数了。
	a := (*int)(unsafe.Pointer(&iniSet1))
	b := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&iniSet1)) + 8))
	fmt.Println(*a, *b)
}


