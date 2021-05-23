package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr [3]uint32
	fmt.Println(uintptr(unsafe.Pointer(&arr)) % 8)
	//对齐边界：取CPU|寄存器字长（机器字长|N bit数据总线) 与数据长度的较小值
	//数据的内存地址：对齐边界的整数倍
	//数据的字节长度：对齐边界的整数倍

	//判断机器字长
	bit := 32 << (^uint(0) >> 63)
	fmt.Println(bit)
}
