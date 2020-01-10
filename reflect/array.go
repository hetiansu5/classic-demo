package main

import (
	"reflect"
	"fmt"
)

func main() {
	var arr [5]interface{}
	arr[0] = 1
	arr[1] = 2
	arr[2] = "3"
	arr[3] = 4
	arr[4] = 5
	t := reflect.TypeOf(arr)
	v := reflect.ValueOf(arr)

	v.Kind()


	//数组的类型，容量，大小，偏移元素值
	fmt.Println(t.Kind(), v.Cap(), v.Len(), v.Index(3))

	for i := 0; i < v.Cap(); i++ {
		//获取数组中元素的具体类型
		tt := reflect.TypeOf(v.Index(i).Interface())
		fmt.Println(tt.Kind())
	}

	var b [5]interface{}
	bv := reflect.ValueOf(&b)
	pbv := reflect.ValueOf(&b)

	fmt.Println(pbv.Elem().Cap())

	for i := 0; i < bv.Elem().Cap(); i++ {
		//反射设置值，非指针不能重置设置值
		fmt.Println(bv.CanSet(), bv.CanAddr())
	}

	for i := 0; i < pbv.Elem().Cap(); i++ {
		//反射设置值，数组不需要初始化容量
		pbv.Elem().Index(i).Set(reflect.ValueOf(i))
	}
	fmt.Println(b)
}
