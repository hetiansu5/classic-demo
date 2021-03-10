package main

import (
	"math"
	"fmt"
)

var matrix []int

func main() {
	n := 9
	matrix = make([]int, n+1)
	i := minDist(n)
	fmt.Println(i)
}

/**
 * 爬楼梯的动态规划
 * min_dist(n) = min(min_dist(n - 1) + min_dist(n - 3) + min_dist(n - 5)) + 1
 */
func minDist(j int) int {
	if j == 0 {
		return 0
	} else if j == 1 || j == 3 || j == 5 {
		return 1
	}

	if matrix[j] > 0 {
		return matrix[j]
	}

	min1, min3, min5 := math.MaxInt64, math.MaxInt64, math.MaxInt64
	if j-5 >= 0 {
		min5 = minDist(j-5) + 1
	}
	if j-3 >= 0 {
		min3 = minDist(j-3) + 1
	}
	if j-1 >= 0 {
		min1 = minDist(j-1) + 1
	}
	min := minValue(min1, min3, min5)
	matrix[j] = min

	return min
}

func minValue(nums ...int) int {
	minV := math.MaxInt64
	for _, v := range nums {
		if minV > v {
			minV = v
		}
	}
	return minV
}
