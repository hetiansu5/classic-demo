package main

import (
	"fmt"
	"text/template"
	"unsafe"
)

// MyTemplate 定义和 template.Template 只是形似
type MyTemplate struct {
	name       string
	parseTree  *unsafe.Pointer
	common     *unsafe.Pointer
}

func main() {
	t := template.New("Foo")
	p := (*MyTemplate)(unsafe.Pointer(t))
	p.name = "Bar" // 关键在这里，突破私有成员
	fmt.Println(p, t)
}
