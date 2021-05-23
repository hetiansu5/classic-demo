package main

import (
	"bytes"
	"fmt"
	"text/template"
	"unsafe"
)

const TEMPLATE = `
template1
{{.Num}}
{{range $key, $value := .Header }}
            Key:{{$key}}, value:{{$value}} <br/>
{{end}}
`

type fire struct {
	Num int
	Header map[string]string
}


func main(){
	map1 := make(map[string]string)
	map1["a"] = "AAA"
	map1["b"] = "BBB"
	map1["c"] = "CCC"
	t, err := template.New("").Parse(TEMPLATE)
	if err != nil {
		fmt.Println(err)
	}
	data := new(fire)
	data.Num = 33
	data.Header = map1
	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf.Bytes()))

	a := unsafe.Sizeof(data.Header)
	fmt.Println(a)

	b := unsafe.Sizeof(map1)
	fmt.Println(b)

	//
	slice := []int{1,3,4, 5}
	c := unsafe.Sizeof(slice)
	fmt.Println(c)
}

