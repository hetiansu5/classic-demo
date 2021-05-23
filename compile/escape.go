package main

func main()  {
	Sum()
}

// Sum 函数返回 0-100 的整数之和
// go run -gcflags '-m -l' escape.go
// go run  -gcflags=-m  escape.go  打印关于编译器的逃逸分析
// output:
//# command-line-arguments
//./escape.go:3:6: can inline main
//./escape.go:11:17: make([]int, count) does not escape
func Sum() int {
	const count = 100
	numbers := make([]int, count)
	for i := range numbers {
		numbers[i] = i + 1
	}

	var sum int
	for _, i := range numbers {
		sum += i
	}
	return sum
}
