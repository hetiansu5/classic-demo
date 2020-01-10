package main

import (
	"testing"
	"strconv"
	"math/rand"
	"fmt"
)

func BenchmarkStrconv(b *testing.B)  {
	var s string
	for i := 0; i < b.N; i++ {
		s = strconv.Itoa(rand.Int())
	}
	s = s + ""
}


func BenchmarkSprint(b *testing.B)  {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint(rand.Int())
	}
	s = s + ""
}

func TestFileWriter_Write(t *testing.T) {
	fmt.Println("222")
}
