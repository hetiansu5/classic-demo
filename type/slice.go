package main

import (
	"strings"
	"fmt"
)

func main() {
	arr := strings.Split("", ",")
	//slice支持游标最大值为len大小
	fmt.Println(arr[1:])
}
