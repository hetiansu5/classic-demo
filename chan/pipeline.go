package main

import "fmt"

// Set up the pipeline.
func main() {
	c := gen(2, 3)
	out := sq(c)
	
	for b := range out {
		fmt.Println(b)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
