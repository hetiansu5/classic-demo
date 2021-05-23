package main

import "fmt"

func main() {
	m := make(map[int]string, 4)
	m[1] = "1"
	m[30] = "30"
	m[2] = "2"
	m[15] = "15"

	for _, v := range m {
		fmt.Println(v)
	}
	fmt.Println("")

	for _, v := range m {
		fmt.Println(v)
	}
	fmt.Println("")

	for _, v := range m {
		fmt.Println(v)
	}
	fmt.Println("")

	for _, v := range m {
		fmt.Println(v)
	}
	fmt.Println("")


}
