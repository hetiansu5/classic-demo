package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//从 slice 中得到一块内存地址,第一个参数为len，第二个参数为cap，只给一个参数表示len和cap一样
	b := make([]byte, 100, 200)
	b[0] = 200
	ptr := unsafe.Pointer(&b[0])
	fmt.Println(len(b), cap(b))
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&b)))

	//从 Go 的内存地址中构造一个 slice
	//var ptr unsafe.Pointer
	var length = 12
	var s1 = struct {
		addr uintptr
		len  int
		cap  int
	}{uintptr(ptr), length, length}
	s := *(*[]byte)(unsafe.Pointer(&s1))
	fmt.Println(s)
	fmt.Println(uintptr(unsafe.Pointer(&s1)))
	fmt.Println(uintptr(unsafe.Pointer(&s)))
	fmt.Println(uintptr(unsafe.Pointer(&b)))
	fmt.Println(uintptr(unsafe.Pointer(&b[0])))

	//通过make初始化的size为0切片，指向元素数组的指针都是同一块内存
	a := make([]int, 0)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&a)))
	a1 := make([]int, 0)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&a1)))
}
