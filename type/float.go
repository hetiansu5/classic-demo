package main

import (
	"strconv"
	"fmt"
)

func main()  {
	var f = float64(2.10100)
	fmt.Println(f)
	fmt.Println(strconv.FormatFloat(f, 'f', -1, 64))
}
