package main

import (
	"reflect"
	"fmt"
)

type PtrChild struct {
	Caption string
}

type Ptr struct {
	Id    int
	Name  string
	Child PtrChild
	Child2 *PtrChild
}

func main() {
	p := &Ptr{}
	//rt := reflect.TypeOf(p)
	rv := reflect.ValueOf(p)
	fmt.Println(rv.Kind(), rv.CanAddr())

	ev := rv.Elem()

	//通过reflect反射给指针变量赋值
	for i := 0; i < ev.NumField(); i++ {
		if ev.CanSet() && ev.CanAddr() {
			if ev.Field(i).Kind() == reflect.Int {
				ev.Field(i).Set(reflect.ValueOf(1))
			} else if ev.Field(i).Kind() == reflect.String {
				ev.Field(i).Set(reflect.ValueOf("33"))
			} else if ev.Field(i).Kind() == reflect.Struct { //嵌套struct
				ptrValue(ev.Field(i))
			} else if ev.Field(i).Kind() == reflect.Ptr { //嵌套指针struct
				fmt.Println(ev.Field(i).IsNil())
				if ev.Field(i).IsNil() {
					fmt.Println(ev.Field(i).Elem().CanSet())
					if ev.Field(i).CanSet() {
						//这里Type()在加Elem()是因为当前Type()是指针，他指向的Elem数据结构才是我们要new的
						ev.Field(i).Set(reflect.New(ev.Field(i).Type().Elem()))
						//函数传递的指向数据的部分携带有指针，可以操作到原数据
						ptrValue(ev.Field(i).Elem())
					}
				}
			}
		}
	}

	fmt.Println(p.Id, p.Name, p.Child.Caption, p.Child2.Caption)

	//case reflect.Ptr:
	//	if out.Type().Elem() == reflect.TypeOf(resolved) {
	//	// TODO DOes this make sense? When is out a Ptr except when decoding a nil value?
	//	elem := reflect.New(out.Type().Elem())
	//	elem.Elem().Set(reflect.ValueOf(resolved))
	//	out.Set(elem)
	//	return true
	//	}
}

func ptrValue(ev reflect.Value)  {
	for i := 0; i < ev.NumField(); i++ {
		fmt.Println(ev.CanSet(), ev.CanAddr())
		if ev.CanSet() && ev.CanAddr() {
			if ev.Field(i).Kind() == reflect.Int {
				ev.Field(i).Set(reflect.ValueOf(22))
			} else if ev.Field(i).Kind() == reflect.String {
				ev.Field(i).Set(reflect.ValueOf("44"))
			}
		}
	}
}
