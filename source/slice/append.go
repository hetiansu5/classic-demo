package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	appendSlice1()
	appendSlice2()
}

// 验证append的母slice的cap容量足够时，append后的新生成slice的底层array指针还是指向原来的内存空间
func appendSlice1() {
	arr1 := make([]int, 5, 11)
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr2 := make([]int, 6, 6)
	arr2[0] = 11
	arr2[1] = 12
	arr2[2] = 13
	arr := append(arr1, arr2...)
	fmt.Println(arr1, len(arr1), cap(arr2))
	fmt.Println(arr2, len(arr2), cap(arr2))
	fmt.Println(arr, len(arr), cap(arr))
	//slice变量的内存地址
	fmt.Println(unsafe.Pointer(&arr1), unsafe.Pointer(&arr2), unsafe.Pointer(&arr))
	//slice底层的array第一个元素的内存地址
	fmt.Println(unsafe.Pointer(&arr1[0]), unsafe.Pointer(&arr2[0]), unsafe.Pointer(&arr[0]))
	reflect.SliceHeader{}

	slice := append([]byte("hello "), "world"...)
	fmt.Println(string(slice))
}

// 验证添加的元素slice的cap容量足够时，并没有用，主要看母slice的cap容量
func appendSlice2() {
	arr1 := make([]int, 5, 5)
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr2 := make([]int, 8, 8)
	arr2[0] = 11
	arr2[1] = 12
	arr2[2] = 13
	arr := append(arr1, arr2...)
	fmt.Println(arr1, len(arr1), cap(arr2))
	fmt.Println(arr2, len(arr2), cap(arr2))
	fmt.Println(arr, len(arr), cap(arr))
	//slice变量的内存地址
	fmt.Println(unsafe.Pointer(&arr1), unsafe.Pointer(&arr2), unsafe.Pointer(&arr))
	//slice底层的array第一个元素的内存地址
	fmt.Println(unsafe.Pointer(&arr1[0]), unsafe.Pointer(&arr2[0]), unsafe.Pointer(&arr[0]))

	slice := append([]byte("hello "), "world"...)
	fmt.Println(string(slice))
}
