package main

import "reflect"

type A interface{}
type B struct{}

var a *B

func main() {
	println(a == nil)            //true
	println(a == (*B)(nil))      //true
	println((A)(a) == (*B)(nil)) //true

	println((A)(a) == nil) //false
	println(reflect.TypeOf((A)(a)).Elem().Kind().String())
}
