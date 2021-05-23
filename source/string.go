package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

type StringHeaderCopy struct {
	Data unsafe.Pointer
	Len  int
}

func main()  {
	var s = "abcdef"

	//string支持slice模式
	fmt.Println(s[2:])

	//slice下标超过最大值
	fmt.Println(s[:6])

	fmt.Println(s[1])

	for k, v := range s {
		fmt.Println(k, v)
	}

	//通过unsafe.Pointer可以转换为普通类型指针
	//这里是为了验证，声明string类型时，底层的stringHeader的Data指针还是nil，并不会开辟字节数据空间
	var m string
	m1 := (*StringHeaderCopy)(unsafe.Pointer(&m))
	fmt.Println(m1)

	var b byte
	b = 'A'  //ok
	//b = '中' //会溢出报错
	fmt.Println(reflect.ValueOf('b').Kind()) //int32
	fmt.Println(reflect.ValueOf(b).Kind()) //uint8
	fmt.Println(reflect.ValueOf(s[1]).Kind()) //uint8
	fmt.Println(s[1] == 'b') //true

	StringJson()
}

type JsonData struct {
	M string `json:"m"`
}

func StringJson(){
	var s = `{"m":"dd\"dd"}` //在字符串中反斜杠，并不是转义字符，真实存在的; json decode后，作为value，才是转义字符
	fmt.Println(string(s[8]), string(s[9]))
	jsonData := JsonData{}
	json.Unmarshal([]byte(s), &jsonData)
	fmt.Println(jsonData.M, string(jsonData.M[2]), string(jsonData.M[3]))

	var s1 = "a1\"b2" //反斜杠是双引号的转移，经过编译后，反斜杠是不存在的
	fmt.Println(string(s1[2]))

	var s2 = `a1\"b2` //反斜杠是真实存在的
	fmt.Println(string(s2[2]), string(s2[3]))
}
