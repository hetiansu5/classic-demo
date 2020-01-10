package main

import (
	"fmt"
)

func main()  {
	var s = "abcdef"

	//string支持slice模式
	fmt.Println(s[2:])

	//slice下标超过最大值
	fmt.Println(s[:6])
}
