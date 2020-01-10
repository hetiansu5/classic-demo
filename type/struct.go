package main

import "fmt"

type Test struct {}

func (t Test) apply(){
	fmt.Println(1)
}

func main() {
	var t Test
	t.apply()
}
