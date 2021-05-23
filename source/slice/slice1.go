package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 4)
	s = []int{1, 2, 3, 4}
	// 指向s变量的指针，当s扩容时，只是改变了其自身Data指针的指向新元素数组
	y := &s
	// 值复制， Data Len Cap这些复制，Data是指向元素数组的指针，当s扩容时，s的Data指针指向新的元素数组，z还是原来的不变
	z := s

	s[0] = 10
	fmt.Println(s, *y, z)

	s = append(s, 1)
	fmt.Println(s, *y, z)

	s[0] = 100
	fmt.Println(s, *y, z)
}
