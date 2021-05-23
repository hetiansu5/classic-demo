package main

import "time"

func main() {
	for {
		println(sum1([]int64{1, 3, 4, 5, 9}))
		time.Sleep(time.Second)
	}
}

func sum1(arr []int64) int64 {
	var j int64
	for _, v := range arr {
		j += v
	}
	return j
}
