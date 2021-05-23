package main

import (
	"fmt"
	"reflect"
)

func main() {
	//map中访问不存在的键，返回值类型的零值
	m := returnNilMap()
	v := m["cdd"]
	fmt.Println(reflect.TypeOf(v))

	var m2 map[string]int
	fmt.Println(m2["d"])
}

func returnNilMap() map[string]string {
	return nil
}
