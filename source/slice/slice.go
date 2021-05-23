package main

import (
	"fmt"
	"strings"
)

func main() {
	//slice支持游标最大值为len大小
	arr := strings.Split("", ",")
	fmt.Println(arr[1:])

	// slice的另类初始化, key表示的是初始化位置，55没有指定就延续上一个元素的索引位置+1
	var iniSet3 = []int{4: 44, 55, 66, 1: 77}
	fmt.Println(iniSet3, len(iniSet3), cap(iniSet3))

	//切片reset，重用
	iniSet3 = iniSet3[0:0]
}

//将slice或者map数据作为参数传递时，需要注意先copy再操作，这是因为slice和map的底层数据结构设计的问题。
func SetTrips(trips []int) {
	tripsCopy := make([]int, len(trips))
	copy(tripsCopy, trips)
	fmt.Println(trips)
}

