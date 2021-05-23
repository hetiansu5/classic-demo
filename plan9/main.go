package main

import (
	"classic-demo/plan9/fab"
	"classic-demo/plan9/register"
	"fmt"
)

func main() {
	a, b, c, d, f := register.SpFp(2)
	fmt.Printf("a:0x%x\nb:0x%x\nc:0x%x\nd:0x%x\nf:0x%x\n", a, b, c, d, f)

	a, b, c, d, f = register.SpZero(2)
	fmt.Printf("a:0x%x\nb:0x%x\nc:0x%x\nd:0x%x\nf:0x%x\n", a, b, c, d, f)

	a1 := fab.Fab(20)
	fmt.Println(a1)
}