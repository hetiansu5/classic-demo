package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

type demoWithOne = struct {
	E int
}
type demoWithTwo = struct {
	E  int
	E1 int
}
type demoWithThree = struct {
	E  int
	E1 int
	E2 int
}

// 测试结构体1指针map
func BenchForStructPointerMapOne(N int) {
	runtime.GC()
	demo1 := make(map[int]*demoWithOne) // KV类型->对应有指针的map
	for i := 0; i < N; i++ {
		v := i
		// 这里分别初始化三种不同的结构体，每种结构体的成员变量类型相同，个数不同
		demo1[i] = &demoWithOne{v}
	}
	// 进行runtime.GC()，并统计时间
	println("结构体1指针开始统计")
	start := time.Now()
	runtime.GC() // 开启这一轮的GC
	result := time.Since(start)
	println(fmt.Sprintf("GC time: %v", result))
}

// 测试结构体2指针map
func BenchForStructPointerMapTwo(N int) {
	runtime.GC()
	demo2 := make(map[int]*demoWithTwo)
	for i := 0; i < N; i++ {
		v := i
		// 这里分别初始化三种不同的结构体，每种结构体的成员变量类型相同，个数不同
		demo2[i] = &demoWithTwo{v, v}
	}
	// 进行runtime.GC()，并统计时间
	println("结构体2指针开始统计")
	start := time.Now()
	runtime.GC() // 开启这一轮的GC
	result := time.Since(start)
	println(fmt.Sprintf("GC time: %v", result))
}

// 测试结构体3指针map
func BenchForStructPointerMapThree(N int) {
	runtime.GC()
	demo3 := make(map[int]*demoWithThree)
	for i := 0; i < N; i++ {
		v := i
		// 这里分别初始化三种不同的结构体，每种结构体的成员变量类型相同，个数不同
		demo3[i] = &demoWithThree{v, v, v}
	}
	// 进行runtime.GC()，并统计时间
	println("结构体3指针开始统计")
	start := time.Now()
	runtime.GC() // 开启这一轮的GC
	result := time.Since(start)
	println(fmt.Sprintf("GC time: %v", result))
}

// 测试指针map
func BenchForPointerMap(N int) {
	runtime.GC()
	demo1 := make(map[int]*int) // KV类型->对应有指针的map
	for i := 0; i < N; i++ {
		v := i
		demo1[i] = &v
	}
	// 进行runtime.GC()，并统计时间
	println("有指针开始统计")
	start := time.Now()
	runtime.GC() // 开启这一轮的GC
	result := time.Since(start)
	println(fmt.Sprintf("GC time: %v", result))
}

// 测试不带指针map
func BenchForNonPointerMap(N int) {
	runtime.GC()
	demo2 := make(map[int]int) // KV类型->对应无指针的map
	for i := 0; i < N; i++ {
		v := i
		demo2[i] = v
	}
	println("无指针开始统计")
	// 进行runtime.GC()，并统计时间
	start := time.Now()
	runtime.GC() // 开启这一轮的GC
	result := time.Since(start)
	println(fmt.Sprintf("Non GC time: %v", result))
}

// -t 1
func main() {
	N := 100000
	var t int
	flag.IntVar(&t, "t", 1, "type")
	flag.Parse()

	switch t {
	case 1:
		BenchForNonPointerMap(N)
		break
	case 2:
		BenchForPointerMap(N)
		break
	case 3:
		BenchForStructPointerMapOne(N)
		break
	case 4:
		BenchForStructPointerMapTwo(N)
		break
	case 5:
		BenchForStructPointerMapThree(N)
	}
}
