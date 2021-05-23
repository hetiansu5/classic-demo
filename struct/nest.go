package main

import (
	"fmt"
)

type Demo1 struct {
}

func (d *Demo1) Do() {
	fmt.Println("demo1 do")
}

func (d *Demo1) Action(){
	fmt.Println("demo1 action")
	d.Do()
}


type Demo2 struct {
	Demo1
}

func (d *Demo2) Do() {
	fmt.Println("demo2 do")
}

func main(){
	demo := &Demo2{}
	demo.Action()
}


