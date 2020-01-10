package main

import (
	"reflect"
	"fmt"
)

type MapInfo struct {
	Map map[string]int8
}

func main() {
	var m map[string]int8
	rv := reflect.ValueOf(m)

	fmt.Println(1 << 5)

	//map结构键的类型
	fmt.Println(rv.Type().Key().Kind())

	//map结构值的类型
	fmt.Println(rv.Type().Elem().Kind())

	//直接反射的map，是不可修改的
	fmt.Println(rv.CanSet(), rv.CanAddr())

	//也无法通过setMapIndex修改或者增加键值
	//rv.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(int8(3)))
	//fmt.Println(rv)


	info := &MapInfo{Map:map[string]int8{
		"c": 1,
	}}
	iv := reflect.ValueOf(info)
	miv := iv.Elem().Field(0)

	mapReflect := reflect.MakeMapWithSize(miv.Type(), 2)


	v1 := reflect.New(miv.Type().Elem())

	mapReflect.SetMapIndex(reflect.ValueOf("a"), v1.Elem())
	mapReflect.SetMapIndex(reflect.ValueOf("b"), v1.Elem())

	fmt.Println(miv.CanSet(), miv.CanAddr(), v1.CanSet())

	miv.SetMapIndex(reflect.ValueOf("d"), v1.Elem())
	for _, k := range miv.MapKeys() {
		fmt.Println(miv.MapIndex(k).CanSet())
	}


	miv.Set(mapReflect)

	fmt.Println(miv)
}
